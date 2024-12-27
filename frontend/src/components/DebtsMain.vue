<template>
    <div class="debt-table">
        <DebtsTable :useActual="useActual" />
        <DebtExport :bills="debts" />
    </div>
</template>

<script lang="ts">
import { defineComponent, computed, onMounted, ref } from "vue";
import { useDebtsStore } from "../store/modules/debts";
import DebtsTable from "./DebtsTable.vue";
import DebtExport from "./DebtExport.vue";

export default defineComponent({
    name: "DebtsMain",
    components: {
        DebtsTable,
        DebtExport,
    },
    setup() {
        const debtsStore = useDebtsStore();

        // Fetch debts on mount
        onMounted(() => {
            debtsStore.fetchDebts();
        });

        const debts = computed(() => debtsStore.debts);
        const useActual = ref(true);

        return {
            debts,
            useActual,
        };
    },
});
</script>
