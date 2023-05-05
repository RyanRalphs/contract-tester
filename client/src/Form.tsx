import React, { useState } from "react";
import { ApiData } from "./types";

type Props = {
  onSubmit: (apiData: ApiData) => boolean;
};

const Form: React.FC<Props> = ({ onSubmit }) => {
  const [url, setUrl] = useState("");
  const [method, setMethod] = useState("");
  const [payload, setPayload] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError("");
    if (isValidUrl(url)) {
      if (payloadIsJson(payload)) {
        fetch("http://localhost:8080/api", {
          method: "POST",
          body: JSON.stringify({ url, method, payload }),
        })
          .then((response) => {
            if (response.status === 200) {
                console.log(response)
                const buildUp = { url, method, payload,outcome: "Passed!" };
                const exists = onSubmit(buildUp); 
                if (exists) {
                    setError("This API Route already exists in the table. Please enter a new API route.");
                    setUrl("");
                    return
                }
                setError("That's a match!");
            }
            if (response.status === 400) {
              setError("What you've submitted as expected is not a match. Please revise your expected payload and submit again.");
              setPayload("");
              return
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