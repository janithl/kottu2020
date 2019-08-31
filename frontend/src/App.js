import React from "react";
import "./App.css";
import { Router } from "@reach/router";

import Page from "./components/Page";

function App() {
  return (
    <div className="App">
      <Router>
        <Page path=":lang/:time" />
        <Page path=":lang" />
        <Page default />
      </Router>
    </div>
  );
}

export default App;
