import React from "react";
import {render} from "react-dom";

function Webapp() {
    return (
      <div>
        <h1>testing</h1>
      </div>
    );
}

render(<Webapp />, document.getElementById("react-target"));
