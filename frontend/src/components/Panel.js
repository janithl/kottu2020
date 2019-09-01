import React from "react";
import { Link } from "@reach/router";

function Panel({ title, subtitle, items, link, linkTitle }) {
  return (
    <div className="panel panel-default">
      <div className="panel-heading">
        <h3 className="panel-title">
          {title} <small>{subtitle}</small>
        </h3>
      </div>
      <div className="list-group">
        {items.map(item => (
          <Link className="list-group-item" to="/">
            <span className="badge">{item.score}</span> {item.title}
          </Link>
        ))}
      </div>
      <div className="panel-footer text-center">
        <Link className="btn btn-success" title={linkTitle} to={link}>
          See More →
        </Link>
      </div>
    </div>
  );
}

export default Panel;
