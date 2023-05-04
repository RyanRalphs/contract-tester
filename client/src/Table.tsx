import React from "react";
import { ApiData } from "./types";
import "./Table.css";

interface Props {
  data: ApiData[];
}

const Table: React.FC<Props> = ({ data }) => {
  return (
    <table className="table mt-3">
      <thead className="thead-dark">
        <tr>
          <th scope="col">API URL</th>
          <th scope="col">Method</th>
          <th scope="col">Expected Outcome</th>
        </tr>
      </thead>
      <tbody>
        {data.map((apiData) => (
          <tr key={apiData.url}>
            <td>{apiData.url}</td>
            <td>{apiData.method}</td>
            <td>{apiData.payload}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default Table;