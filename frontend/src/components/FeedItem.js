import React from "react";

function FeedItem({ item }) {
  return (
    <article className="panel panel-default">
      <div className="panel-body">
        <h3>
          <a href="/go/?id={item.ID}" id="{item.ID}">
            {item.Title}
          </a>
        </h3>
        <a href="/blog/{item.BlogID}">{item.BlogID}</a>

        <div className="media hidden-xs">
          <div className="media-left media-top">
            <img className="media-object" alt="alt" src="link" />
          </div>

          <div className="media-body">{item.Content}</div>
        </div>
      </div>

      <div className="panel-footer">
        <button className="btn btn-link btn-xs btn-ts">
          {item.CreatedAtRemote}
        </button>
        <button
          className="btn btn-primary btn-xs btn-facebook"
          title="This post was liked/shared {item.ShareCount} time(s)"
        >
          <span className="glyphicon glyphicon-share-alt"></span>
          {" Shares: " + item.ShareCount}
        </button>
        <button
          className="btn btn-link btn-xs pull-right"
          title={"The post popularity is " + item.Chilies + " chilies"}
        >
          <img
            src={"/icons/chili" + item.Chilies + ".png"}
            alt={item.Chilies + " chilies"}
          />
        </button>
      </div>
    </article>
  );
}

export default FeedItem;
