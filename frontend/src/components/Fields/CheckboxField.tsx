import 'bulma/css/bulma.min.css';
import { Form } from 'react-bulma-components';

type CheckboxFieldCfg = {
  idx: number
  checked: boolean[],
  setChecked: React.Dispatch<React.SetStateAction<boolean[]>>
}

function CheckboxField({ idx, checked, setChecked }: CheckboxFieldCfg) {
  return (
    <Form.Checkbox
      checked={checked[idx]}
      onChange={(e) => {
        setChecked((c) => c.map((el, i) => idx == i ? e.target.checked : el))
      }}>
    </Form.Checkbox>
  )
}

export default CheckboxField;
