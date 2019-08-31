import React from "react";

import Header from "./Header";
import Content from "./Content";
import Feed from "./Feed";

function Page({ lang, time }) {
  return (
    <>
      <Header lang={lang} time={time} />
      <Content>
        <Feed endpoint={"http://localhost:9000/api/latest/" + lang} />
      </Content>
    </>
  );
}

Page.defaultProps = {
  lang: "en",
  time: "off"
};

export default Page;
