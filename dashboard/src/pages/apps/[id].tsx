import { useParams } from '@solidjs/router'
import { createResource } from 'solid-js'
import { client } from '/@/libs/api'
import { Header } from '/@/components/Header'
import { container, contentContainer } from '/@/pages/apps.css'
import { applicationState, getWebsiteURL, providerToIcon } from '/@/libs/application'
import {
  appTitle,
  card,
  cardItem,
  cardItemContent,
  cardItems,
  cardItemTitle,
  cardTitle,
  centerInline,
} from '/@/pages/apps/[id].css'
import { StatusIcon } from '/@/components/StatusIcon'
import { titleCase } from '/@/libs/casing'
import { Application_ContainerState, DeployType } from '/@/api/neoshowcase/protobuf/apiserver_pb'
import { durationHuman } from '/@/libs/format'

export default () => {
  const params = useParams()
  const [app] = createResource(
    () => params.id,
    (id) => client.getApplication({ id }),
  )

  return (
    <div className={container}>
      <Header />
      <div className={appTitle}>
        <div className={centerInline}>{providerToIcon('GitHub', 36)}</div>
        <div>{app()?.name}</div>
      </div>
      <div className={contentContainer}>
        <div className={card}>
          <div className={cardTitle}>Overall</div>
          <div className={cardItems}>
            <div className={cardItem}>
              <div className={cardItemTitle}>状態</div>
              <div className={cardItemContent}>
                {app() && <StatusIcon state={applicationState(app())} />}
                {app() && applicationState(app())}
              </div>
            </div>
            {app() && app().deployType === DeployType.RUNTIME && (
              <div className={cardItem}>
                <div className={cardItemTitle}>コンテナの状態</div>
                <div className={cardItemContent}>{app() && titleCase(Application_ContainerState[app().container])}</div>
              </div>
            )}
            <div className={cardItem}>
              <div className={cardItemTitle}>起動時間</div>
              <div className={cardItemContent}>
                {app()?.running && durationHuman(3 * 60 * 1000 /* TODO: use updated_at */)}
                {app() && !app().running && '-'}
              </div>
            </div>
            <div className={cardItem}>
              <div className={cardItemTitle}>作成日</div>
              <div className={cardItemContent}>{app() && new Date().toLocaleString() /* TODO: use created_at */}</div>
            </div>
            {app() && app().websites.length > 0 && (
              <div className={cardItem}>
                <div className={cardItemTitle}>URLs</div>
                <div className={cardItemContent}>
                  {app()?.websites.map((website) => (
                    <a href={getWebsiteURL(website)} target='_blank' rel="noreferrer">
                      {getWebsiteURL(website)}
                    </a>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
