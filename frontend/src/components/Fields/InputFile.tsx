import 'bulma/css/bulma.min.css';
import { Form } from 'react-bulma-components';


type InputFieldCfg = {
  lable: string,
  placeholder: string,
  name: string,

  onChange(e: React.ChangeEvent<HTMLInputElement>): void,

  error?: string
}

function InputFile({
  lable,
  placeholder,
  onChange,
  error
 }: InputFieldCfg ) {
  return (
    <Form.Field>
      <Form.Label>{lable}</Form.Label>
      <Form.Control>
        <Form.InputFile
          fullwidth
          align='right'
          label={placeholder}
          onChange={onChange}
          color={error && "danger"}
        />
        {error && <Form.Help color="danger">{error}</Form.Help>}
      </Form.Control>
    </Form.Field>
  )
}

export default InputFile;
