import { A, useParams } from '@solidjs/router'
import { createResource, For, onCleanup, Show } from 'solid-js'
import { client } from '/@/libs/api'
import { Header } from '/@/components/Header'
import { applicationState, buildTypeStr, getWebsiteURL, providerToIcon } from '/@/libs/application'
import { StatusIcon } from '/@/components/StatusIcon'
import { titleCase } from '/@/libs/casing'
import { Application_ContainerState, BuildConfig, DeployType } from '/@/api/neoshowcase/protobuf/gateway_pb'
import { DiffHuman, shortSha } from '/@/libs/format'
import { vars } from '/@/theme'
import { styled } from '@macaron-css/solid'
import { Container } from '/@/libs/layout'
import { URLText } from '/@/components/URLText'
import { Button } from '/@/components/Button'

const AppTitleContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    gap: '14px',
    alignContent: 'center',

    marginTop: '48px',
    fontSize: '32px',
    fontWeight: 'bold',
    color: vars.text.black1,
  },
})

const AppTitle = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    gap: '8px',
  },
})

const CenterInline = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
  },
})

const CardsContainer = styled('div', {
  base: {
    marginTop: '24px',
    display: 'flex',
    flexDirection: 'row',
    flexWrap: 'wrap',
    gap: '40px',
  },
})

const Card = styled('div', {
  base: {
    minWidth: '320px',
    borderRadius: '4px',
    border: `1px solid ${vars.bg.white4}`,
    background: vars.bg.white1,
    padding: '24px 36px',

    display: 'flex',
    flexDirection: 'column',
    gap: '24px',
  },
})

const CardTitle = styled('div', {
  base: {
    fontSize: '24px',
    fontWeight: 600,
  },
})

const CardItems = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '20px',
  },
})

const CardItem = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    gap: '8px',
  },
})

const CardItemTitle = styled('div', {
  base: {
    fontSize: '16px',
    color: vars.text.black3,
  },
})

const CardItemContent = styled('div', {
  base: {
    marginLeft: 'auto',
    fontSize: '16px',
    color: vars.text.black1,

    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    gap: '4px',
  },
})

interface BuildConfigInfoProps {
  config: BuildConfig
}

const BuildConfigInfo = (props: BuildConfigInfoProps) => {
  const c = props.config.buildConfig
  switch (c.case) {
    case 'runtimeCmd':
      return (
        <>
          <CardItem>
            <CardItemTitle>Base Image</CardItemTitle>
            <CardItemContent>{c.value.baseImage || 'Scratch'}</CardItemContent>
          </CardItem>
          {c.value.buildCmd && (
            <CardItem>
              <CardItemTitle>Build Command{c.value.buildCmdShell && ' (Shell)'}</CardItemTitle>
              <CardItemContent>{c.value.buildCmd}</CardItemContent>
            </CardItem>
          )}
        </>
      )
    case 'runtimeDockerfile':
      return (
        <>
          <CardItem>
            <CardItemTitle>Dockerfile</CardItemTitle>
            <CardItemContent>{c.value.dockerfileName}</CardItemContent>
          </CardItem>
        </>
      )
    case 'staticCmd':
      return (
        <>
          <CardItem>
            <CardItemTitle>Base Image</CardItemTitle>
            <CardItemContent>{c.value.baseImage || 'Scratch'}</CardItemContent>
          </CardItem>
          {c.value.buildCmd && (
            <CardItem>
              <CardItemTitle>Build Command{c.value.buildCmdShell && ' (Shell)'}</CardItemTitle>
              <CardItemContent>{c.value.buildCmd}</CardItemContent>
            </CardItem>
          )}
          <CardItem>
            <CardItemTitle>Artifact Path</CardItemTitle>
            <CardItemContent>{c.value.artifactPath}</CardItemContent>
          </CardItem>
        </>
      )
    case 'staticDockerfile':
      return (
        <>
          <CardItem>
            <CardItemTitle>Dockerfile</CardItemTitle>
            <div>{c.value.dockerfileName}</div>
          </CardItem>
          <CardItem>
            <CardItemTitle>Artifact Path</CardItemTitle>
            <div>{c.value.artifactPath}</div>
          </CardItem>
        </>
      )
  }
}

export default () => {
  const params = useParams()
  const [app, { refetch: refetchApp }] = createResource(
    () => params.id,
    (id) => client.getApplication({ id }),
  )
  const [repo] = createResource(
    () => app()?.repositoryId,
    (id) => client.getRepository({ repositoryId: id }),
  )
  const loaded = () => !!(app() && repo())

  const refetchTimer = setInterval(refetchApp, 10000)
  onCleanup(() => clearInterval(refetchTimer))

  const startApp = async () => {
    await client.startApplication({ id: app().id })
    await refetchApp()
  }
  const stopApp = async () => {
    await client.stopApplication({ id: app().id })
    await refetchApp()
  }

  return (
    <Container>
      <Header />
      <Show when={loaded()}>
        <AppTitleContainer>
          <CenterInline>{providerToIcon('GitHub', 36)}</CenterInline>
          <AppTitle>
            <div>{repo().name}</div>
            <div>/</div>
            <div>{app().name}</div>
          </AppTitle>
        </AppTitleContainer>
        <CardsContainer>
          <Card>
            <Show when={!app().running}>
              <Button color='black1' size='large' onclick={startApp}>
                Start App
              </Button>
            </Show>
            <Show when={app().running}>
              <Button color='black1' size='large' onclick={startApp}>
                Restart App
              </Button>
            </Show>
            <Show when={app().running}>
              <Button color='black1' size='large' onclick={stopApp}>
                Stop App
              </Button>
            </Show>
          </Card>
          <Card>
            <CardTitle>Overall</CardTitle>
            <CardItems>
              <CardItem>
                <CardItemTitle>状態</CardItemTitle>
                <CardItemContent>
                  <StatusIcon state={applicationState(app())} size={24} />
                  {applicationState(app())}
                </CardItemContent>
              </CardItem>
              <Show when={app().deployType === DeployType.RUNTIME}>
                <CardItem>
                  <CardItemTitle>コンテナの状態</CardItemTitle>
                  <CardItemContent>{app() && titleCase(Application_ContainerState[app().container])}</CardItemContent>
                </CardItem>
              </Show>
              <CardItem>
                <CardItemTitle>起動時刻</CardItemTitle>
                <CardItemContent>
                  <Show when={app().running} fallback={'-'}>
                    <DiffHuman target={app().updatedAt.toDate()} />
                  </Show>
                </CardItemContent>
              </CardItem>
              <CardItem>
                <CardItemTitle>作成日</CardItemTitle>
                <CardItemContent>
                  <DiffHuman target={app().createdAt.toDate()} />
                </CardItemContent>
              </CardItem>
              <Show when={app().websites.length > 0}>
                <CardItem>
                  <CardItemTitle>URLs</CardItemTitle>
                </CardItem>
                <CardItem>
                  <CardItemTitle>
                    <For each={app().websites} children={(website) => (
                      <URLText href={getWebsiteURL(website)} target='_blank' rel="noreferrer">
                        {getWebsiteURL(website)}
                      </URLText>
                    )} />
                  </CardItemTitle>
                </CardItem>
              </Show>
            </CardItems>
          </Card>
          <Card>
            <CardTitle>Info</CardTitle>
            <CardItems>
              <CardItem>
                <CardItemTitle>ID</CardItemTitle>
                <CardItemContent>{app().id}</CardItemContent>
              </CardItem>
              <CardItem>
                <CardItemTitle>Name</CardItemTitle>
                <CardItemContent>{app().name}</CardItemContent>
              </CardItem>
              <A href={`/repos/${repo().id}`}>
                <CardItem>
                  <CardItemTitle>Repository</CardItemTitle>
                  <CardItemContent>
                    <CenterInline>{providerToIcon('GitHub', 20)}</CenterInline>
                    {repo().name}
                  </CardItemContent>
                </CardItem>
              </A>
              <CardItem>
                <CardItemTitle>Git ref (short)</CardItemTitle>
                <CardItemContent>{app().refName}</CardItemContent>
              </CardItem>
              <CardItem>
                <CardItemTitle>Deploy type</CardItemTitle>
                <CardItemContent>{titleCase(DeployType[app().deployType])}</CardItemContent>
              </CardItem>
              <CardItem>
                <CardItemTitle>Commit</CardItemTitle>
                <CardItemContent>
                  <Show
                    when={app().currentCommit !== app().wantCommit}
                    fallback={<div>{shortSha(app().currentCommit)}</div>}
                  >
                    <div>
                      {shortSha(app().currentCommit)} → {shortSha(app().wantCommit)} (Deploying)
                    </div>
                  </Show>
                </CardItemContent>
              </CardItem>
              <CardItem>
                <CardItemTitle>Use MariaDB</CardItemTitle>
                <CardItemContent>{`${app().config.useMariadb}`}</CardItemContent>
              </CardItem>
              <CardItem>
                <CardItemTitle>Use MongoDB</CardItemTitle>
                <CardItemContent>{`${app().config.useMongodb}`}</CardItemContent>
              </CardItem>
            </CardItems>
          </Card>
          <Card>
            <CardTitle>Build Config</CardTitle>
            <CardItems>
              <CardItem>
                <CardItemTitle>Build Type</CardItemTitle>
                <CardItemContent>{buildTypeStr[app().config.buildConfig.buildConfig.case]}</CardItemContent>
              </CardItem>A
              <BuildConfigInfo config={app().config.buildConfig} />
              <Show when={app().config.entrypoint}>
                <CardItem>
                  <CardItemTitle>Entrypoint</CardItemTitle>
                  <CardItemContent>{app()?.config.entrypoint}</CardItemContent>
                </CardItem>
              </Show>
              <Show when={app().config.command}>
                <CardItem>
                  <CardItemTitle>Command</CardItemTitle>
                  <CardItemContent>{app()?.config.command}</CardItemContent>
                </CardItem>
              </Show>
            </CardItems>
          </Card>
        </CardsContainer>
      </Show>
    </Container>
  )
}
