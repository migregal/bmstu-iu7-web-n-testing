import { Dispatch, SetStateAction } from "react";

import { Modal } from "react-bulma-components";

import UploadModelForm from "components/Forms/UploadModelForm";

function UploadModel(
  { modal, setModal }: { modal: boolean, setModal: Dispatch<SetStateAction<boolean>> }
) {
  return (
    <Modal show={modal} onClose={() => { setModal(false) }}>
      <Modal.Card>
        <Modal.Card.Header>
          <Modal.Card.Title>
            Upload Model
          </Modal.Card.Title>
        </Modal.Card.Header>
        <Modal.Card.Body>
          <UploadModelForm setInProgress={setModal} />
        </Modal.Card.Body>
        <Modal.Card.Footer>
        </Modal.Card.Footer>
      </Modal.Card >
    </Modal >
  )
}

export default UploadModel;
