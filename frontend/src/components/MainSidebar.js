import React from "react";

import Panel from "./Panel";

function MainSidebar() {
  return (
    <div className="sidebar col-sm-12 col-md-4">
      <Panel
        title="Trending On Kottu"
        subtitle="Today"
        items={[]}
        link="/all/trending"
      />
      <Panel title="Hot Posts" subtitle="Today" items={[]} link="/all/hot" />
    </div>
  );
}

export default MainSidebar;
