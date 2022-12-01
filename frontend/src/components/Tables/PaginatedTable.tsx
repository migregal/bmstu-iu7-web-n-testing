import { NavigateOptions, URLSearchParamsInit } from 'react-router-dom';

import 'bulma/css/bulma.min.css';
import { Heading, Pagination, Table } from 'react-bulma-components';


import { PAGE_QUERY_PARAM } from 'contstants/pagination';

type PaginatedTableParams<T> = {
  head: React.ReactElement[],
  data: T[],
  total: number,
  map(d: T, i: number): React.ReactElement,
  params: URLSearchParams,
  setParams: (nextInit?: URLSearchParamsInit | ((prev: URLSearchParams) => URLSearchParamsInit), navigateOpts?: NavigateOptions) => void,
  onPageChanged: () => void
}

function PaginatedTable<T>(
  { head, data, total, map, params, setParams, onPageChanged}: PaginatedTableParams<T>
) {
  return !data ?
    (<Heading renderAs='h2' style={{ textAlign: 'center' }}>
      Seems like there is no data yet...
    </Heading>) :
    (<div>
      <Table.Container>
        <Table bordered size="fullwidth">
          <thead>
            <tr>
              {head.map((h, i) => <td key={i}>{h}</td>)}
            </tr>
          </thead>
          <tbody>
            {data.map((d, i) => map(d, i))}
          </tbody>
        </Table>
      </Table.Container>
      <Pagination
        autoHide
        rounded
        showFirstLast
        current={Number(params.get(PAGE_QUERY_PARAM)) ?? 1}
        total={total}
        onChange={(current: number) => {
          params.set(PAGE_QUERY_PARAM, current.toString());
          setParams(params, { replace: true });
          onPageChanged();
        }}
      />
    </div >
    )
}

export default PaginatedTable;
