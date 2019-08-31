import React from "react";
import { Link } from "@reach/router";

const languages = {
  en: "English",
  si: "සිංහල",
  ta: "தமிழ்"
};

function LangSelector({ lang, time }) {
  return (
    <ul className="nav navbar-nav navbar-right">
      {Object.keys(languages).map(l => (
        <li key={l} className={lang === l ? "langs active" : "menuitem"}>
          <Link className="menuitem" to={"/" + l}>
            {languages[l]}
          </Link>
        </li>
      ))}
    </ul>
  );
}

LangSelector.defaultProps = {
  lang: "en",
  time: "off"
};

export default LangSelector;
