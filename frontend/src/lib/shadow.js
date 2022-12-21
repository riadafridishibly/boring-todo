import Editor from "./Editor.svelte";
import cssData from "bytemd/dist/index.css";

export function append() {
  const target = document.createElement("div");
  const root = target.attachShadow({ mode: "open" });
  const style = document.createElement("style");

  root.appendChild(style);

  style.type = "text/css";
  style.appendChild(document.createTextNode(cssData));

  new Editor({
    target: root,
    props: {},
  });

  document.body.append(target);
}
