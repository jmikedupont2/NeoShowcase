import LogoImage from '/@/assets/logo.svg?url'
import { user } from '/@/libs/api'
import { colorVars } from '/@/theme'
import { styled } from '@macaron-css/solid'
import { A } from '@solidjs/router'
import { Component, Show } from 'solid-js'
import { UserMenuButton } from '../UI/UserMenuButton'

const Container = styled('div', {
  base: {
    width: '100%',
    height: '64px',
    padding: '10px 24px',
    flexShrink: 0,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    borderBottom: `1px solid ${colorVars.semantic.ui.border}`,
  },
})

export const Header: Component = () => {
  return (
    <Container>
      <A href="/">
        <img src={LogoImage} alt="NeoShowcase logo" />
      </A>
      <Show when={user()}>{(user) => <UserMenuButton user={user()} />}</Show>
    </Container>
  )
}
