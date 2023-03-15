<template>
  <div class="home">
    <DebtsMain />
    <NotesSection />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import NotesSection from "./NotesSection.vue";
import DebtsMain from "./DebtsMain.vue";

export default defineComponent({
  name: "Home",
  components: {DebtsMain, NotesSection},
  props: {},
  setup() {
    return {};
  },
});
</script>

<style>

</style>