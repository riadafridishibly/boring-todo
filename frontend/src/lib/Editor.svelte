<script>
  import styles from "bytemd/dist/index.css";
  import { Editor } from "bytemd";
  import gfm from "@bytemd/plugin-gfm";

  const plugins = [gfm()];
  let value = "# Hello world\n\nWelcome to the editor!\n";

  function handleChange(e) {
    value = e.detail.value;
  }
</script>

<Editor {value} {plugins} on:change={handleChange} />
