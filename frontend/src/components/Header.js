import React from "react";

import LangSelector from "./LangSelector";
import SearchBar from "./SearchBar";

function Header({ lang, time }) {
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
            <a className="navbar-brand" href="./">
              <img
                alt="Kottu"
                title="Go to the Kottu home page"
                src="/logo.png"
              />
            </a>
          </div>
          <div className="collapse navbar-collapse" id="navbar-collapse">
            <p className="navbar-text hidden-sm">
              <strong className="visible-xs-inline">Kottu </strong>syndicates
              over{" "}
              <a title="View our blogroll" href="/blogroll">
                100 Sri Lankan blogs
              </a>
              . You can{" "}
              <a
                title="Learn more about Kottu and how you can join"
                href="./about"
              >
                join too
              </a>
              .
            </p>
            <SearchBar />
            <LangSelector lang={lang} time={time} />
          </div>
        </div>
      </nav>
    </header>
  );
}

Header.defaultProps = {
  lang: "en",
  time: "off"
};

export default Header;
