import React from "react";

import Header from "./Header";
import Content from "./Content";
import MainSidebar from "./MainSidebar";
import { LatestFeed } from "./Feed";

function Page({ lang, time }) {
  return (
    <>
      <Header lang={lang} time={time} />
      <div className="container">
        <Content>
          <LatestFeed lang={lang} />
        </Content>
        <MainSidebar />
      </div>
    </>
  );
}

Page.defaultProps = {
  lang: "en",
  time: "off"
};

export default Page;
