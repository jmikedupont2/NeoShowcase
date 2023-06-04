import { Header } from '/@/components/Header'
import { Checkbox } from '/@/components/Checkbox'
import { createResource, createSignal, For, Show } from 'solid-js'
import { Radio, RadioItem } from '/@/components/Radio'
import { client, user } from '/@/libs/api'
import {
  Application,
  GetApplicationsRequest_Scope,
  GetRepositoriesRequest_Scope,
} from '/@/api/neoshowcase/protobuf/gateway_pb'
import { RepositoryRow } from '/@/components/RepositoryRow'
import { ApplicationState } from '/@/libs/application'
import { StatusCheckbox } from '/@/components/StatusCheckbox'
import { styled } from '@macaron-css/solid'
import { vars } from '/@/theme'
import { Container } from '/@/libs/layout'
import { Button } from '/@/components/Button'
import { A } from '@solidjs/router'

const sortItems: RadioItem<string>[] = [
  { value: 'desc', title: '最新順' },
  { value: 'asc', title: '古い順' },
]

const scopeItems = (admin: boolean) => {
  const items: RadioItem<GetRepositoriesRequest_Scope>[] = [
    { value: GetRepositoriesRequest_Scope.MINE, title: '自分のアプリ' },
    { value: GetRepositoriesRequest_Scope.PUBLIC, title: 'すべてのアプリ' },
  ]
  if (admin) {
    items.push({ value: GetRepositoriesRequest_Scope.ALL, title: 'すべてのアプリ (admin)' })
  }
  return items
}

const AppsTitle = styled('div', {
  base: {
    marginTop: '48px',
    fontSize: '32px',
    fontWeight: 'bold',
    color: vars.text.black1,
  },
})

const ContentContainer = styled('div', {
  base: {
    marginTop: '24px',
    display: 'grid',
    gridTemplateColumns: '380px 1fr',
    gap: '40px',
  },
})

const SidebarContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '22px',

    padding: '24px 40px',
    backgroundColor: vars.bg.white1,
    borderRadius: '4px',
    border: `1px solid ${vars.bg.white4}`,
  },
})

const SidebarSection = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '16px',
  },
})

const SidebarTitle = styled('div', {
  base: {
    fontSize: '24px',
    fontWeight: 500,
    color: vars.text.black1,
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

const MainContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '20px',
  },
})

const SearchBarContainer = styled('div', {
  base: {
    display: 'grid',
    gridTemplateColumns: '1fr 180px',
    gap: '20px',
    height: '44px',
  },
})

const SearchBar = styled('input', {
  base: {
    padding: '12px 20px',
    borderRadius: '4px',
    border: `1px solid ${vars.bg.white4}`,
    fontSize: '14px',

    '::placeholder': {
      color: vars.text.black3,
    },
  },
})

const RepositoriesContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '20px',
  },
})

export default () => {
  const [scope, setScope] = createSignal(GetRepositoriesRequest_Scope.MINE)
  const appScope = () => {
    const mine = scope() === GetRepositoriesRequest_Scope.MINE
    return mine ? GetApplicationsRequest_Scope.MINE : GetApplicationsRequest_Scope.ALL
  }

  const [repos] = createResource(
    () => scope(),
    (scope) => client.getRepositories({ scope }),
  )
  const [apps] = createResource(
    () => appScope(),
    (scope) => client.getApplications({ scope }),
  )
  const loaded = () => !!(user() && repos() && apps())

  const appsByRepo = () =>
    loaded() &&
    apps().applications.reduce((acc, app) => {
      if (!acc[app.repositoryId]) acc[app.repositoryId] = []
      acc[app.repositoryId].push(app)
      return acc
    }, {} as Record<string, Application[]>)

  const [sort, setSort] = createSignal(sortItems[0].value)

  return (
    <Container>
      <Header />
      <AppsTitle>Apps</AppsTitle>
      <Show when={loaded()}>
        <ContentContainer>
          <SidebarContainer>
            <SidebarSection>
              <SidebarTitle>Status</SidebarTitle>
              <SidebarOptions>
                <Checkbox>
                  <StatusCheckbox apps={apps().applications} state={ApplicationState.Idle} title='Idle' />
                </Checkbox>
                <Checkbox>
                  <StatusCheckbox apps={apps().applications} state={ApplicationState.Deploying} title='Deploying' />
                </Checkbox>
                <Checkbox>
                  <StatusCheckbox apps={apps().applications} state={ApplicationState.Running} title='Running' />
                </Checkbox>
                <Checkbox>
                  <StatusCheckbox apps={apps().applications} state={ApplicationState.Static} title='Static' />
                </Checkbox>
              </SidebarOptions>
            </SidebarSection>
            <SidebarSection>
              <SidebarTitle>Scope</SidebarTitle>
              <SidebarOptions>
                <Radio items={scopeItems(user().admin)} selected={scope()} setSelected={setScope} />
              </SidebarOptions>
            </SidebarSection>
            <SidebarOptions>
              <SidebarTitle>Sort</SidebarTitle>
              <Radio items={sortItems} selected={sort()} setSelected={setSort} />
            </SidebarOptions>
          </SidebarContainer>
          <MainContainer>
            <SearchBarContainer>
              <SearchBar placeholder='Search...' />
              <A href='/repos/new'>
                <Button color='black1' size='large'>
                  + New Repository
                </Button>
              </A>
            </SearchBarContainer>
            <RepositoriesContainer>
              <For each={repos().repositories}>{(r) => <RepositoryRow repo={r} apps={appsByRepo()[r.id] || []} />}</For>
            </RepositoriesContainer>
          </MainContainer>
        </ContentContainer>
      </Show>
    </Container>
  )
}
