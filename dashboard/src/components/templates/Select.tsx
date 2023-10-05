import DropDownIcon from '/@/assets/icons/24/expand_more.svg'
import { clickInside as clickInsideDir, clickOutside as clickOutsideDir } from '/@/libs/useClickInout'
import { colorVars, textVars } from '/@/theme'
import { style } from '@macaron-css/core'
import { styled } from '@macaron-css/solid'
import { For, JSX, Show, createSignal } from 'solid-js'
import { Button } from '../UI/Button'
import { CheckBoxIcon } from '../UI/CheckBoxIcon'

// https://github.com/solidjs/solid/discussions/845
const clickInside = clickInsideDir
const clickOutside = clickOutsideDir

const Container = styled('div', {
  base: {
    position: 'relative',
    width: '100%',
    display: 'flex',
    flexDirection: 'column',
    gap: '4px',
  },
})
const SelectButton = styled('button', {
  base: {
    width: '100%',
    height: '48px',
    overflowX: 'auto',
    padding: '10px 16px',
    display: 'grid',
    gridTemplateColumns: '1fr 24px',
    alignItems: 'center',
    gap: '4px',

    background: colorVars.semantic.ui.primary,
    borderRadius: '8px',
    border: 'none',
    outline: `1px solid ${colorVars.semantic.ui.border}`,
    color: colorVars.semantic.text.black,
    cursor: 'pointer',

    selectors: {
      '&:focus': {
        outline: `2px solid ${colorVars.semantic.primary.main}`,
      },
      '&:disabled': {
        cursor: 'not-allowed',
        color: colorVars.semantic.text.disabled,
        background: colorVars.semantic.ui.tertiary,
      },
    },
  },
  variants: {
    opened: {
      true: {
        outline: `2px solid ${colorVars.semantic.primary.main}`,
      },
    },
  },
})
const Title = styled('div', {
  base: {
    width: '100%',
    ...textVars.text.regular,
    whiteSpace: 'nowrap',
    overflow: 'hidden',
    textOverflow: 'ellipsis',
    textAlign: 'left',
  },
  variants: {
    placeholder: {
      true: {
        color: colorVars.semantic.text.disabled,
      },
    },
  },
})
const DropDownIconContainer = styled('div', {
  base: {
    width: '24px',
    height: '24px',
    flexShrink: 0,
  },
})
const optionsContainerClass = style({
  position: 'absolute',
  width: '100%',
  maxHeight: '250px',
  top: '56px',
  overflowY: 'auto',
  padding: '6px',

  display: 'flex',
  flexDirection: 'column',

  background: colorVars.semantic.ui.primary,
  borderRadius: '6px',
  boxShadow: '0px 0px 20px 0px rgba(0, 0, 0, 0.10)',
  zIndex: 1,
})

export interface SelectItem<T> {
  value: T
  title: string
}

export type SingleSelectProps<T> = {
  items: SelectItem<T>[]
  setSelected: (s: T) => void
  disabled?: boolean
} & (
  | {
      selected: T
      placeHolder?: string
    }
  | {
      selected: undefined
      placeHolder: string
    }
)

export const SingleSelect = <T,>(props: SingleSelectProps<T>): JSX.Element => {
  const [showOptions, setShowOptions] = createSignal(false)

  const showPlaceHolder = () => props.selected === undefined
  const selectedTitle = () => props.items.find((i) => i.value === props.selected)?.title ?? props.placeHolder

  const handleSelect = (item: SelectItem<T>) => {
    props.setSelected(item.value)
    setShowOptions(false)
  }

  return (
    <Container>
      <SelectButton
        onClick={() => {
          setShowOptions((s) => !s)
        }}
        opened={showOptions()}
        disabled={props.disabled}
        type="button"
      >
        <Title placeholder={showPlaceHolder()}>{selectedTitle()}</Title>
        <DropDownIconContainer>
          <DropDownIcon />
        </DropDownIconContainer>
      </SelectButton>
      {/* TODO: help text */}
      <Show when={showOptions()}>
        <div
          use:clickInside={() => setShowOptions(true)}
          use:clickOutside={() => setShowOptions(false)}
          class={optionsContainerClass}
        >
          <For each={props.items}>
            {(item) => (
              <Button color="text" size="medium" full onClick={() => handleSelect(item)}>
                {item.title}
              </Button>
            )}
          </For>
        </div>
      </Show>
    </Container>
  )
}

export type MultiSelectProps<T> = {
  items: SelectItem<T>[]
  setSelected: (s: T[]) => void
  disabled?: boolean
} & (
  | {
      selected: T[]
      placeHolder?: string
    }
  | {
      selected: undefined
      placeHolder: string
    }
)

export const MultiSelect = <T,>(props: MultiSelectProps<T>): JSX.Element => {
  const [showOptions, setShowOptions] = createSignal(false)
  const showPlaceHolder = () => props.selected === undefined || props.selected?.length === 0

  const selectedTitle = () => {
    if (showPlaceHolder()) {
      return props.placeHolder
    } else {
      return props.selected?.map((s) => props.items.find((i) => i.value === s)?.title).join(', ')
    }
  }

  const handleSelect = (item: SelectItem<T>) => {
    if (props.selected?.includes(item.value)) {
      props.setSelected(props.selected.filter((s) => s !== item.value))
    } else {
      props.setSelected([...(props.selected ?? []), item.value])
    }
  }

  return (
    <Container>
      <SelectButton
        onClick={() => {
          setShowOptions((s) => !s)
        }}
        opened={showOptions()}
        disabled={props.disabled}
      >
        <Title placeholder={showPlaceHolder()}>{selectedTitle()}</Title>
        <DropDownIconContainer>
          <DropDownIcon />
        </DropDownIconContainer>
      </SelectButton>
      {/* TODO: help text */}
      <Show when={showOptions()}>
        <div
          use:clickInside={() => setShowOptions(true)}
          use:clickOutside={() => setShowOptions(false)}
          class={optionsContainerClass}
        >
          <For each={props.items}>
            {(item) => {
              const checked = () => props.selected?.includes(item.value) ?? false
              return (
                <Button
                  color="text"
                  size="medium"
                  hasCheckbox
                  full
                  onClick={() => handleSelect(item)}
                  leftIcon={<CheckBoxIcon checked={checked()} />}
                >
                  {item.title}
                </Button>
              )
            }}
          </For>
        </div>
      </Show>
    </Container>
  )
}
