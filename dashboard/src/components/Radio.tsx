import { vars } from '/@/theme'
import { styled } from '@macaron-css/solid'
import { ImRadioChecked, ImRadioUnchecked } from 'solid-icons/im'
import { For, JSXElement } from 'solid-js'

const Container = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'column',
    gap: '12px',

    fontSize: '20px',
    color: vars.text.black1,
  },
})

const ItemContainer = styled('div', {
  base: {
    display: 'flex',
    flexDirection: 'row',
    gap: '12px',

    cursor: 'pointer',
    alignItems: 'center',
  },
})

export interface RadioItem<T extends string | number> {
  value: T
  title: string
}

export interface Props<T extends string | number> {
  items: RadioItem<T>[]
  selected: T
  setSelected: (s: T) => void
  onClick?: () => void
}

export const Radio = <T extends string | number>(props: Props<T>): JSXElement => {
  return (
    <Container>
      <For each={props.items}>
        {(item) => (
          <ItemContainer
            onClick={() => {
              props.setSelected(item.value)
              props.onClick?.()
            }}
          >
            {props.selected === item.value ? (
              <ImRadioChecked size={20} color={vars.text.black2} />
            ) : (
              <ImRadioUnchecked size={20} color={vars.text.black4} />
            )}
            {item.title}
          </ItemContainer>
        )}
      </For>
    </Container>
  )
}
