import { Dispatch, SetStateAction, useEffect, useState } from "react";

import { Modal, Progress } from "react-bulma-components";

import EditModelForm from "components/Forms/EditModelForm";
import { useWeightsStore } from "contexts/weightsContext";
import { toast } from "react-toastify";

interface EditModelProps {
  modal: boolean
  setModal: Dispatch<SetStateAction<boolean>>
  modelID: string
  structureID: string
}

function EditModel({ modal, setModal, modelID, structureID }: EditModelProps) {
  const weights = useWeightsStore();

  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (!modal)
      return;

    weights.getByStructureID(structureID)
      .catch((err) => toast(err.message))
      .finally(() => setIsLoading(false))
  }, [modal]);

  return (
    <Modal show={modal} onClose={() => { setModal(false) }}>
      {isLoading ? <Progress color="primary" /> :
        <Modal.Card>
          <Modal.Card.Header>
            <Modal.Card.Title>
              Edit Model
            </Modal.Card.Title>
          </Modal.Card.Header>
          <Modal.Card.Body>
            <EditModelForm
              modelID={modelID}
              setInProgress={setModal}
            />
          </Modal.Card.Body>
          <Modal.Card.Footer>
          </Modal.Card.Footer>
        </Modal.Card >
      }
    </Modal >
  )
}

export default EditModel;
