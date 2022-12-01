import { useEffect, useState } from "react";
import { useSearchParams } from "react-router-dom";

import { Observer } from "mobx-react";

import { Box, Button, Container, Form, Heading, Level, Progress } from "react-bulma-components";

import { toast } from "react-toastify";

import { PAGE_QUERY_PARAM, PER_PAGE_QUERY_PARAM } from "contstants/pagination";

import { useUsersStore } from "contexts/usersContext";
import { useAuthStore } from "contexts/authContext";

import Page from "components/Pages/Page";
import PaginatedTable from 'components/Tables/PaginatedTable';
import CheckboxField from "components/Fields/CheckboxField";
import CheckboxAllField from "components/Fields/CheckboxAllField";

function Users() {
  const auth = useAuthStore();
  const users = useUsersStore();

  const [searchParams, setSeatchParams] = useSearchParams()

  const [isLoading, setIsLoading] = useState(false);

  const [checked, setChecked] = useState<boolean[]>([]);

  const getPerPage = () => Number(searchParams.get(PER_PAGE_QUERY_PARAM)) || 10

  const getTotalPages = () => Math.ceil(users.users.total / getPerPage())

  const getPage = () => Math.min(Number(searchParams.get(PAGE_QUERY_PARAM)) - 1 || 0, getTotalPages())

  const getUsers = () => {
    setIsLoading(true)
    users.getAll({ page: getPage(), perPage: getPerPage() })
      .then(() => setChecked(() => new Array(users.users.infos?.length ?? 0).fill(false)))
      .catch((err) => toast(err.message))
      .finally(() => setIsLoading(false));
  }

  useEffect(() => getUsers(), []);

  const anyChecked = () => {
    return auth.isadmin() && checked.filter((el) => el == true).length != 0
  }

  const [blockDate, setBlockDate] = useState(new Date());

  return (
    <Page>
      <Container
        className="contentContainer"
        style={{ width: isLoading ? "50vw" : "100vw" }}
      >
        <Box>
          <Level>
            <Level.Side>
              <Heading>Users</Heading>
            </Level.Side>
            {anyChecked() && <Level.Side align="right">
              <Level.Item>
                <Button.Group>
                  <Button
                    outlined
                    color="link"
                    loading={isLoading}
                    onClick={() => {
                      setIsLoading(() => true);
                      users
                        .unblock({ ids: users.users.infos.filter((u, i) => checked[i]).map((u) => u.id) })
                        .catch((err) => toast(err.message))
                        .finally(() => setIsLoading(false));
                    }}
                  >Unblock</Button>
                </Button.Group>
              </Level.Item>
              <Level.Item>
                <Form.Input type={"datetime-local"} onChange={(e) => { setBlockDate(() => new Date(e.target.value)) }} />
              </Level.Item>
              <Level.Item>
                <Button.Group>
                  <Button
                    outlined
                    color="danger"
                    loading={isLoading}
                    onClick={() => {
                      setIsLoading(() => true);
                      users
                        .block({ ids: users.users.infos.filter((u, i) => checked[i]).map((u) => u.id), until: blockDate })
                        .catch((err) => toast(err.message))
                        .finally(() => setIsLoading(false));
                    }}
                  >Block</Button>
                </Button.Group>
              </Level.Item>
            </Level.Side>}
          </Level>
          {isLoading ? <Progress color="primary" /> :
            <Observer>
              {() =>
                <PaginatedTable
                  head={
                    (!auth.isadmin() ? [] : [<CheckboxAllField checked={checked} setChecked={setChecked} />])
                      .concat([<>Username</>, <>Email</>, <>Fullname</>])
                      .concat(auth.isadmin() ? [<>Block</>] : [])
                  }
                  data={users.users.infos}
                  total={users.users.total}
                  map={(d, i) => (
                    <tr key={d.username}>
                      {auth.isadmin() && <td align="center">
                        <CheckboxField idx={i} checked={checked} setChecked={setChecked} />
                      </td>}
                      <td>{d.username}</td>
                      <td>{d.email}</td>
                      <td>{d.fullname}</td>
                      {auth.isadmin() && <td>{d.block}</td>}
                    </tr>)}
                  params={searchParams}
                  setParams={setSeatchParams}
                  onPageChanged={getUsers}
                />
              }
            </Observer>
          }
        </Box >
      </Container >
    </Page >
  );
}

export default Users;
