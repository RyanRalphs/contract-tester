import React, { useState } from "react";
import "./App.css";
import { ApiData } from "./types";
import Form from "./Form";
import Table from "./Table";

const App: React.FC = () => {
  const [apiData, setApiData] = useState<ApiData[]>([]);

  const handleApiDataSubmit = (apiDataItem: ApiData)  => {
    let exists = false
    apiData.forEach((item) => {
      if (item.url === apiDataItem.url) {
        exists = true
        return exists
      }})

      if (exists) {
        return exists
      }

        setApiData((prevApiData) => [...prevApiData, apiDataItem])
        return exists


  };

  return (
    <div className="container">
      <h1 className="text-center mt-3">API Form</h1>
      <Form onSubmit={handleApiDataSubmit} />

      {apiData.length > 0 && <Table data={apiData} />}
    </div>
  );
};

export default App;