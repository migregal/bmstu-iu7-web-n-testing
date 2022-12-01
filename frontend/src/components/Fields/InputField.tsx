import { HTMLInputTypeAttribute } from 'react';

import 'bulma/css/bulma.min.css';
import { Form, Icon } from 'react-bulma-components';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faExclamationTriangle, IconDefinition } from '@fortawesome/free-solid-svg-icons'

type InputFieldCfg = {
  lable: string,
  placeholder?: string,
  name: string,
  type?: HTMLInputTypeAttribute

  onChange(e: React.ChangeEvent<HTMLInputElement>): void,

  icon?: IconDefinition
  rightIcon?: IconDefinition
  onClick?(e: React.MouseEvent<HTMLSpanElement>): void,

  value?: string
  error?: string
}

function InputField({
  lable,
  placeholder,
  name,
  type,
  onChange,
  icon,
  rightIcon,
  onClick,
  value,
  error
}: InputFieldCfg) {
  return (
    <Form.Field>
      <Form.Label>{lable}</Form.Label>
      <Form.Control>
        <Form.Input
          placeholder={placeholder}
          name={name}
          value={value}
          onChange={onChange}
          type={type}
          color={error && "danger"} />
        {error && <Form.Help color="danger">{error}</Form.Help>}
        {icon && <Icon align="left"><FontAwesomeIcon icon={icon} /></Icon>}
        {error && !rightIcon && <Icon align="right">
          <FontAwesomeIcon icon={faExclamationTriangle} />
        </Icon>}
        {rightIcon && <Icon
          align="right"
          style={{ pointerEvents: "auto" }}
          onClick={onClick}>
          <FontAwesomeIcon icon={rightIcon} />
        </Icon>}
      </Form.Control>
    </Form.Field>
  )
}

export default InputField;
