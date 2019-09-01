import React from "react";
import { Link } from "@reach/router";

import { APIContext } from "../contexts/APIContext";
import LangSelector from "./LangSelector";
import SearchBar from "./SearchBar";

function Introduction({ blogCount }) {
  return (
    <p className="navbar-text hidden-sm">
      <strong className="visible-xs-inline">Kottu </strong>syndicates over{" "}
      <Link title="View our blogroll" to="/blogroll">
        {blogCount} Sri Lankan blogs
      </Link>
      . You can{" "}
      <Link title="Learn more about Kottu and how you can join" to="/about">
        join too
      </Link>
      .
    </p>
  );
}

function Header({ lang, time, numBlogs }) {
  return (
    <header>
      <nav id="mainmenu" className="navbar navbar-inverse navbar-fixed-top">
        <div className="container">
          <div className="navbar-header">
            <button
              type="button"
              className="navbar-toggle collapsed"
              data-toggle="collapse"
              data-target="#navbar-collapse"
              aria-expanded="false"
            >
              <span className="sr-only">Toggle navigation</span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
            </button>
            <Link className="navbar-brand" to="/">
              <img
                alt="Kottu"
                title="Go to the Kottu home page"
                src="/logo.png"
              />
            </Link>
          </div>
          <div className="collapse navbar-collapse" id="navbar-collapse">
            <Introduction blogCount={numBlogs} />
            <SearchBar />
            <LangSelector lang={lang} time={time} />
          </div>
        </div>
      </nav>
    </header>
  );
}

const HeaderWithContext = props => (
  <APIContext.Consumer>
    {api => <Header numBlogs={api.getNumBlogs()} {...props} />}
  </APIContext.Consumer>
);

export default HeaderWithContext;
