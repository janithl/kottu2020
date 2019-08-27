import React from "react";

function Header() {
  return (
    <header>
      <nav id="mainmenu" class="navbar navbar-inverse navbar-fixed-top">
        <div class="container">
          <div class="navbar-header">
            <button
              type="button"
              class="navbar-toggle collapsed"
              data-toggle="collapse"
              data-target="#navbar-collapse"
              aria-expanded="false"
            >
              <span class="sr-only">Toggle navigation</span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
              <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="./">
              <img
                alt="Kottu"
                title="Go to the Kottu home page"
                src="logo.png"
              />
            </a>
          </div>
          <div class="collapse navbar-collapse" id="navbar-collapse">
            <p class="navbar-text hidden-sm">
              <strong class="visible-xs-inline">Kottu </strong>syndicates over{" "}
              <a title="View our blogroll" href="./blogroll">
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
            <form
              class="navbar-form navbar-left"
              method="GET"
              action="./search/"
            >
              <div class="form-group">
                <input
                  tabindex="1"
                  id="searchbar"
                  name="q"
                  type="text"
                  class="form-control"
                  placeholder="Search Kottu..."
                  value=""
                />
              </div>
              <button
                tabindex="2"
                id="searchbtn"
                type="submit"
                class="btn btn-default"
              >
                Search
              </button>
            </form>
            <ul class="nav navbar-nav navbar-right">
              <li class="langs">
                <a class="menuitem" title="View English language posts" href="">
                  English
                </a>
              </li>
              <li class="langs">
                <a class="menuitem" title="View Sinhala language posts" href="">
                  සිංහල
                </a>
              </li>
              <li class="langs">
                <a class="menuitem" title="View Tamil language posts" href="">
                  தமிழ்
                </a>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    </header>
  );
}

export default Header;
