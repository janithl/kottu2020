import React, { useState, useEffect } from "react";

import FeedItem from "./FeedItem";

function Feed({ endpoint }) {
  const [items, setItems] = useState([]);
  const [url, setURL] = useState(endpoint);

  useEffect(() => {
    if (items.length === 0 || url !== endpoint) {
      fetch(endpoint)
        .then(response => response.json())
        .then(json => {
          setItems(json);
          setURL(endpoint);
        })
        .catch(error => console.error(error));
    }
  });

  return (
    <>
      {items.map(item => (
        <FeedItem key={item.ID} item={item} />
      ))}
    </>
  );
}

export default Feed;
