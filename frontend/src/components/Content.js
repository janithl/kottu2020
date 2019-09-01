import React from "react";

function Content({ children }) {
  return (
    <div className="col-sm-12 col-md-8">
      <div className="content">{children}</div>
    </div>
  );
}

export default Content;
