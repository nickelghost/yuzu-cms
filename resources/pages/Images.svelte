<script>
  import TopBar from '../components/TopBar.svelte';

  let fileInput;
  let images = [];
  let isDraggingOver = false;

  function uploadImages(files) {
    if (files.length === 0) return;
    // TODO: Add a check for the field being set back to empty
    // TODO: Make a request to the API
    alert('Image uploaded');
  }
</script>

<style>
  .upload-area {
    align-items: center;
    border: 1px solid #ececec;
    display: flex;
    height: 30vh;
    justify-content: center;
    padding: 18px;
  }

  .upload-area.is-active {
    border: 1px solid green;
  }

  .images {
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;
    padding: 18px 0;
  }

  .image {
    border: 1px solid #ececec;
    height: 144px;
    margin-bottom: 12px;
    width: 256px;
  }
</style>

<TopBar title="Images" />
<main class="content">
  <!-- TODO: Only allow image file formats -->
  <input
    bind:this="{fileInput}"
    multiple="true"
    style="display:none;"
    type="file"
    on:change="{(e) => uploadImages(e.target.files)}"
  />
  <div
    class="upload-area"
    class:is-active="{isDraggingOver}"
    ondragover="return false"
    on:click="{() => fileInput.click()}"
    on:drop|preventDefault="{(e) => {
      isDraggingOver = false;
      uploadImages(e.dataTransfer.files);
    }}"
    on:dragenter|preventDefault="{() => {
      isDraggingOver = true;
    }}"
    on:dragleave|preventDefault="{() => {
      isDraggingOver = false;
    }}"
  >
    Upload images
  </div>
  <div class="images">
    {#each images as image}
      <img alt="custom user uploaded image" class="image" src="{image.src}" />
    {/each}
  </div>
</main>
