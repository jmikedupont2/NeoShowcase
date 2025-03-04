import {
  CreateRepositoryAuth,
  Repository_AuthMethod,
  UpdateRepositoryRequest,
  User,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { FormTextBig } from '/@/components/AppsNew'
import { Button } from '/@/components/Button'
import { Header } from '/@/components/Header'
import { InfoTooltip } from '/@/components/InfoTooltip'
import { InputBar, InputLabel } from '/@/components/Input'
import { RepositoryAuthSettings } from '/@/components/RepositoryAuthSettings'
import RepositoryNav from '/@/components/RepositoryNav'
import { UserSearch } from '/@/components/UserSearch'
import { client, handleAPIError } from '/@/libs/api'
import { extractRepositoryNameFromURL } from '/@/libs/application'
import { Container } from '/@/libs/layout'
import { userFromId, users } from '/@/libs/useAllUsers'
import useModal from '/@/libs/useModal'
import { vars } from '/@/theme'
import { PartialMessage, PlainMessage } from '@bufbuild/protobuf'
import { styled } from '@macaron-css/solid'
import { useParams } from '@solidjs/router'
import { Component, JSX, Show, createEffect, createMemo, createResource, createSignal } from 'solid-js'
import { createStore } from 'solid-js/store'
import toast from 'solid-toast'

const ContentContainer = styled('div', {
  base: {
    marginTop: '24px',
    display: 'grid',
    gridTemplateColumns: '380px 1fr',
    gap: '40px',
    position: 'relative',
  },
})
const SidebarContainer = styled('div', {
  base: {
    position: 'sticky',
    top: '64px',
    padding: '24px 40px',
    backgroundColor: vars.bg.white1,
    borderRadius: '4px',
    border: `1px solid ${vars.bg.white4}`,
  },
})
const SidebarOptions = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '12px',

    fontSize: '20px',
    color: vars.text.black1,
  },
})
const SidebarNavAnchor = styled('a', {
  base: {
    color: vars.text.black2,
    textDecoration: 'none',
    selectors: {
      '&:hover': {
        color: vars.text.black1,
      },
    },
  },
})
const ConfigsContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '24px',
  },
})
const SettingFieldSet = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '16px',
    padding: '24px',
    border: `1px solid ${vars.bg.white4}`,
    borderRadius: '4px',
    background: vars.bg.white1,
  },
})

export default () => {
  const params = useParams()
  const [repo, { refetch: refetchRepo }] = createResource(
    () => params.id,
    (repositoryId) => client.getRepository({ repositoryId }),
  )
  const loaded = () => !!(users() && repo())

  const update = async (req: PlainMessage<UpdateRepositoryRequest>) => {
    try {
      await client.updateRepository(req)
      toast.success('リポジトリ設定を更新しました')
      refetchRepo()
    } catch (e) {
      handleAPIError(e, 'リポジトリ設定の更新に失敗しました')
    }
  }

  const GeneralConfigsContainer: Component = () => {
    let formContainer: HTMLFormElement

    // 現在の設定で初期化
    const [generalConfig, setGeneralConfig] = createStore<PlainMessage<UpdateRepositoryRequest>>({
      id: repo().id,
      name: repo().name,
      url: repo().url,
    })
    const [updateAuthConfig, setUpdateAuthConfig] = createSignal(false)
    const mapAuthMethod = (authMethod: Repository_AuthMethod): PlainMessage<CreateRepositoryAuth>['auth'] => {
      switch (authMethod) {
        case Repository_AuthMethod.NONE:
          return { case: 'none', value: {} }
        case Repository_AuthMethod.BASIC:
          return { case: 'basic', value: { username: '', password: '' } }
        case Repository_AuthMethod.SSH:
          return { case: 'ssh', value: { keyId: '' } }
      }
    }
    const [authConfig, setAuthConfig] = createStore<PlainMessage<CreateRepositoryAuth>>({
      auth: mapAuthMethod(repo().authMethod),
    })

    // URLからリポジトリ名を自動入力
    createEffect(() => {
      const repositoryName = extractRepositoryNameFromURL(generalConfig.url)
      setGeneralConfig('name', repositoryName)
    })

    const onClickSave: JSX.EventHandler<HTMLButtonElement, MouseEvent> = async (e) => {
      e.preventDefault()
      if (!formContainer.reportValidity()) {
        return
      }
      return update({ ...generalConfig, auth: authConfig })
    }

    return (
      <form ref={formContainer}>
        <SettingFieldSet>
          <FormTextBig id="general-settings">General Settings</FormTextBig>
          <div>
            <InputLabel>Repository Name</InputLabel>
            <InputBar
              placeholder="my-app"
              value={generalConfig.name}
              onChange={(e) => setGeneralConfig('name', e.currentTarget.value)}
              required
            />
          </div>
          <div>
            <InputLabel>Repository URL</InputLabel>
            <InputBar
              // SSH URLはURLとしては不正なのでtypeを変更
              type={authConfig.auth.case === 'ssh' ? 'text' : 'url'}
              placeholder="https://example.com/my-app.git"
              value={generalConfig.url}
              onChange={(e) => setGeneralConfig('url', e.currentTarget.value)}
              required
            />
          </div>
          <Show
            when={!updateAuthConfig()}
            fallback={
              <Button color="black1" size="large" width="auto" onclick={() => setUpdateAuthConfig(false)}>
                認証方法の更新をキャンセルする
              </Button>
            }
          >
            <Button color="black1" size="large" width="auto" onclick={() => setUpdateAuthConfig(true)}>
              認証方法を更新する
            </Button>
          </Show>
          <Show when={updateAuthConfig()}>
            <RepositoryAuthSettings authConfig={authConfig} setAuthConfig={setAuthConfig} />
          </Show>
          <Button color="black1" size="large" width="auto" onclick={onClickSave} type="submit">
            Save
          </Button>
        </SettingFieldSet>
      </form>
    )
  }

  const OwnerConfigContainer: Component = () => {
    const { Modal, open } = useModal()

    const nonOwnerUsers = createMemo(() => {
      return users().filter((user) => !repo().ownerIds.includes(user.id)) ?? []
    })

    const handleAddOwner = async (user: User) => {
      const updateApplicationRequest: PartialMessage<UpdateRepositoryRequest> = {
        id: repo().id,
        ownerIds: {
          ownerIds: repo().ownerIds.concat(user.id),
        },
      }

      try {
        await client.updateRepository(updateApplicationRequest)
        toast.success('リポジトリオーナーを追加しました')
        refetchRepo()
      } catch (e) {
        handleAPIError(e, 'リポジトリオーナーの追加に失敗しました')
      }
    }
    const handleDeleteOwner = async (owner: User) => {
      const updateApplicationRequest: PartialMessage<UpdateRepositoryRequest> = {
        id: repo().id,
        ownerIds: {
          ownerIds: repo().ownerIds.filter((id) => id !== owner.id),
        },
      }

      try {
        await client.updateRepository(updateApplicationRequest)
        toast.success('リポジトリのオーナーを削除しました')
        refetchRepo()
      } catch (e) {
        handleAPIError(e, 'リポジトリのオーナーの削除に失敗しました')
      }
    }

    return (
      <>
        <SettingFieldSet>
          <FormTextBig id="owner-settings">
            Owners
            <InfoTooltip
              tooltip={['オーナーは以下が可能になります', 'リポジトリの設定を変更']}
              style="bullets-with-title"
            />
          </FormTextBig>
          <Button color="black1" size="large" width="auto" onclick={open}>
            リポジトリオーナーを追加する
          </Button>
          <UserSearch users={repo().ownerIds.map((userId) => userFromId(userId))}>
            {(user) => (
              <Button
                color="black1"
                size="large"
                width="auto"
                onclick={() => {
                  handleDeleteOwner(user)
                }}
              >
                削除
              </Button>
            )}
          </UserSearch>
        </SettingFieldSet>
        <Modal>
          <UserSearch users={nonOwnerUsers()}>
            {(user) => (
              <Button
                color="black1"
                size="large"
                width="auto"
                onclick={() => {
                  handleAddOwner(user)
                }}
              >
                追加
              </Button>
            )}
          </UserSearch>
        </Modal>
      </>
    )
  }

  return (
    <Container>
      <Header />
      <Show when={loaded()}>
        <RepositoryNav repository={repo()} />
        <ContentContainer>
          <div>
            <SidebarContainer>
              <SidebarOptions>
                <SidebarNavAnchor href="#general-settings">General</SidebarNavAnchor>
                <SidebarNavAnchor href="#owner-settings">Owner</SidebarNavAnchor>
              </SidebarOptions>
            </SidebarContainer>
          </div>
          <ConfigsContainer>
            <GeneralConfigsContainer />
            <OwnerConfigContainer />
          </ConfigsContainer>
        </ContentContainer>
      </Show>
    </Container>
  )
}
