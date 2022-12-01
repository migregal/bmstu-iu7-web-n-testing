import 'bulma/css/bulma.min.css';
import { Form } from 'react-bulma-components';

type CheckboxAllFieldCfg = {
  checked: boolean[],
  setChecked: React.Dispatch<React.SetStateAction<boolean[]>>
}

function CheckboxAllField({ checked, setChecked }: CheckboxAllFieldCfg) {
  return (
    <div style={{ textAlign: "center" }}>
      <Form.Checkbox
        checked={checked.every((c) => c)}
        onChange={(e) => setChecked((c) => c.map(() => e.target.checked))}
      ></Form.Checkbox>
    </div>
  )
}

export default CheckboxAllField;
