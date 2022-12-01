import { useEffect, useState } from "react";
import { useSearchParams } from "react-router-dom";

import { Observer } from "mobx-react";

import { Box, Button, Container, Heading, Level, Progress } from "react-bulma-components";

import { toast } from "react-toastify";

import { PAGE_QUERY_PARAM, PER_PAGE_QUERY_PARAM } from "contstants/pagination";

import { useAuthStore } from "contexts/authContext";
import { useModelsStore } from "contexts/modelsContext";

import Page from "components/Pages/Page";
import UploadModel from "components/Modals/UploadModel";
import EditModel from "components/Modals/EditModel";
import PaginatedTable from 'components/Tables/PaginatedTable';
import CheckboxField from "components/Fields/CheckboxField";
import CheckboxAllField from "components/Fields/CheckboxAllField";

function Models() {
  const auth = useAuthStore();
  const models = useModelsStore();

  const [searchParams, setSearchParams] = useSearchParams()

  const [isLoading, setIsLoading] = useState(false);

  const [checked, setChecked] = useState<boolean[]>([])

  const getPerPage = () => Number(searchParams.get(PER_PAGE_QUERY_PARAM)) || 10

  const getTotalPages = () => Math.ceil(models.models.total / getPerPage())

  const getPage = () => Math.min(Number(searchParams.get(PAGE_QUERY_PARAM)) - 1 || 0, getTotalPages())

  const getModels = () => {
    setIsLoading(true)

    models.getAll({ page: getPage(), perPage: getPerPage() })
      .then(() => setChecked(() => new Array(models.models.infos?.length ?? 0).fill(false)))
      .catch((err) => toast(err.message))
      .finally(() => setIsLoading(false))
  }

  useEffect(() => getModels(), []);

  const anyChecked = () => {
    return checked.filter((el) => el == true).length != 0
  }

  const canEdit = () => {
    return checked.filter((el) => el == true).length == 1
      && models.models.infos.at(checked.indexOf(true))?.owner_id == auth.id()
  }

  const canDownload = () => {
    return anyChecked()
  }

  const canDelete = () => {
    return anyChecked()
      && (auth.isadmin()
        || models.models.infos.filter((el, i) => checked[i]).every((m) => m.owner_id == auth.id()))
  }

  const [addModal, setAddModal] = useState(false);
  const [editModal, setEditModal] = useState(false);

  return (
    <Page>
      <Container
        className="contentContainer"
        style={{ width: isLoading ? "50vw" : "100vw" }}
      >
        <UploadModel modal={addModal} setModal={setAddModal} />

        {canEdit() && <EditModel
          modal={editModal}
          setModal={setEditModal}
          structureID={models.models.infos.at(checked.indexOf(true))?.structure.id ?? ""}
          modelID={models.models.infos.at(checked.indexOf(true))?.id ?? ""}
        />}

        <Box>
          <Level>
            <Level.Side>
              <Heading>Models</Heading>
            </Level.Side>
            <Level.Side align="right">
              <Button.Group>
                {anyChecked() || <Button color="primary" onClick={() => setAddModal(() => true)}>Add</Button>}
                {canEdit() && <Button outlined color="info" onClick={() => setEditModal(() => true)}>Edit</Button>}
                {canDownload() &&
                  <Button
                    outlined
                    color="link"
                    loading={isLoading}
                    onClick={() => {
                      setIsLoading(() => true);
                      models
                        .download({ ids: models.models.infos.filter((m, i) => checked[i]).map((m) => m.id) })
                        .catch((err) => toast(err.message))
                        .finally(() => setIsLoading(() => false));
                    }}>Download</Button>
                }
                {canDelete() &&
                  <Button
                    outlined
                    color="danger"
                    onClick={() => {
                      setIsLoading(() => true);
                      models
                        .delete({ ids: models.models.infos.filter((m, i) => checked[i]).map((m) => m.id) })
                        .catch((err) => toast(err.message))
                        .finally(() => setIsLoading(() => false));
                    }}>Delete</Button>
                }
              </Button.Group>
            </Level.Side>
          </Level>

          {isLoading ? <Progress color="primary" /> :
            <Observer>
              {() =>
                <PaginatedTable
                  head={[
                    <CheckboxAllField checked={checked} setChecked={setChecked} />,
                    <>Owner</>,
                    <>Title</>
                  ]}
                  data={models.models.infos}
                  total={getTotalPages()}
                  map={(d, i) => (
                    <tr key={d.title}>
                      <td align="center">
                        <CheckboxField idx={i} checked={checked} setChecked={setChecked} />
                      </td>
                      <td>{d.owner_name || d.owner_id}</td>
                      <td>{d.title}</td>
                    </tr>)}
                  params={searchParams}
                  setParams={setSearchParams}
                  onPageChanged={getModels}
                />
              }
            </Observer>
          }
        </Box >
      </Container >
    </Page>
  );
}

export default Models;
