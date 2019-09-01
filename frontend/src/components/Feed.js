import React from "react";

import { APIContext } from "../contexts/APIContext";
import FeedItem from "./FeedItem";

export default function Feed({ items }) {
  return (
    <>{items && items.map(item => <FeedItem key={item.ID} item={item} />)}</>
  );
}

export function LatestFeed({ lang }) {
  return (
    <APIContext.Consumer>
      {api => <Feed items={api.getLatestPosts(lang)} />}
    </APIContext.Consumer>
  );
}
