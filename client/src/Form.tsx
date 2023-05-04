import React, { useState } from "react";
import { ApiData } from "./types";

type Props = {
  onSubmit: (apiData: ApiData) => void;
};

const Form: React.FC<Props> = ({ onSubmit }) => {
  const [url, setUrl] = useState("");
  const [method, setMethod] = useState("");
  const [payload, setPayload] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (isValidUrl(url)) {
      if (payloadIsJson(payload)) {
        const buildUp = { url, method, payload };
        onSubmit(buildUp); 
        fetch("http://localhost:8080/api", {
          method: "POST",
          body: JSON.stringify({ url, method, payload }),
        })
          .then((response) => {
            if (response.status === 200) {
              if (response.body) return response.status;
            }
          })
          .then((data) => console.log(data));
      } else {
        setError("Please enter a valid JSON payload");
        setPayload("");
      }
    } else {
      setError("Please enter a valid URL");
      setUrl("");
    }
  };

  const isValidUrl = (url: string) => {
    try {
      new URL(url);
      return true;
    } catch {
      return false;
    }
  };

  const payloadIsJson = (payload: string) => {
    try {
      JSON.parse(payload);
      return true;
    } catch {
      return false;
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="form-group">
        <label>API URL:</label>
        <input
          type="text"
          className="form-control"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="https://example.com/api"
          required
        />
      </div>
      <div className="form-group">
        <label>Expected Outcome:</label>
        <input
          type="text"
          className="form-control"
          onChange={(e) => setPayload(e.target.value)}
          placeholder='{"name": "John Doe"}'
          required
        />
      </div>
      <div className="form-group">
        <label>Method:</label>
        <select
          className="form-control"
          value={method}
          onChange={(e) => setMethod(e.target.value)}
          required
        >
          <option value="">Select a method...</option>
          <option value="GET">GET</option>
          <option value="POST">POST</option>
          <option value="PUT">PUT</option>
          <option value="DELETE">DELETE</option>
        </select>
      </div>
      <button type="submit" className="btn btn-primary">
        Submit
      </button>
      {error && <div className="error">{error}</div>}
    </form>
  );
};

export default Form;