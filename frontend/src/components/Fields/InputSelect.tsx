import 'bulma/css/bulma.min.css';
import { Form } from 'react-bulma-components';


type InputFieldCfg = {
  lable: string,
  placeholder: string,

  children: JSX.Element
  onChange(e: React.ChangeEvent<HTMLSelectElement>): void,

  error?: string
}

function InputSelect({ lable, onChange, children }: InputFieldCfg) {
  return (
    <Form.Field>
      <Form.Control>
        <Form.Label>{lable}</Form.Label>
        <Form.Select onChange={onChange}>
          <option value={""}>
            Add a new one
          </option>
          {children}
        </Form.Select>
      </Form.Control>
    </Form.Field>
  )
}

export default InputSelect;
