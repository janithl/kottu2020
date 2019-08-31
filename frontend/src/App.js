import React from "react";
import "./App.css";
import { Router } from "@reach/router";

import Header from "./components/Header";

function App() {
  return (
    <div className="App">
      <Router>
        <Header path=":lang/:time" />
        <Header path=":lang" />
        <Header default />
      </Router>
    </div>
  );
}

export default App;
