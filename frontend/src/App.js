import React, { useState } from "react";
import { Router } from "@reach/router";
import "./App.css";

import { APIContext, API } from "./contexts/APIContext";
import Page from "./components/Page";

function App() {
  const [apiState, setApiState] = useState({});

  const getState = key => apiState[key];
  const setState = (key, value) => setApiState({ [key]: value });

  const url = "http://localhost:9000/api/";
  const api = new API(url, getState, setState);

  return (
    <APIContext.Provider value={api}>
      <div className="App">
        <Router>
          <Page path=":lang/:time" />
          <Page path=":lang" />
          <Page default />
        </Router>
      </div>
    </APIContext.Provider>
  );
}

export default App;
