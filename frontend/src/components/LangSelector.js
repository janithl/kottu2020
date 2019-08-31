import React from "react";

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
          <a className="menuitem" href={"/" + l}>
            {languages[l]}
          </a>
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
