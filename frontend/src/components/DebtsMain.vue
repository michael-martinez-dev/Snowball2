<template>
    <div class="debt-table">
        <DebtsTable :useActual="useActual" />
        <v-row>
            <v-col cols="10">
                <v-text-field
                    v-if="changeStoreDir"
                    v-model="selectedDirectory"
                    :loading="saving"
                    density="compact"
                    placeholder="Storage location"
                    :rules="[directoryPathRule]"
                    clearable
                    autofocus
                    hide-details
                    variant="solo"
                    @keydown.enter="handleDirSelection"
                    @keydown.esc="
                        () => {
                            changeStoreDir = false;
                            saving = false;
                        }
                    "
                    @blur="
                        () => {
                            changeStoreDir = false;
                            saving = false;
                        }
                    "
                />
                <v-btn
                    v-else
                    density="compact"
                    @click="
                        () => {
                            changeStoreDir = true;
                        }
                    "
                >
                    Change Storage Location
                </v-btn>
            </v-col>
            <v-col>
                <DebtExport :bills="debts" />
            </v-col>
        </v-row>
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
        const selectedDirectory = ref<String | null>(null);
        const windowsDirectoryRegex = new RegExp(
            '^[A-Za-z]:\\\\(?:[^<>:"|?*\\\\\r\n]+\\\\)*(?:[^<>:"|?*\\\\\r\n]+)?$',
        );
        const changeStoreDir = ref(false);
        const saving = ref(false);

        // Fetch debts on mount
        onMounted(() => {
            debtsStore.fetchDebts();
        });

        const debts = computed(() => debtsStore.debts);
        const useActual = ref(true);

        const handleDirSelection = () => {
            if (
                selectedDirectory.value === null ||
                selectedDirectory.value === ""
            ) {
                changeStoreDir.value = false;
                saving.value = false;
                return;
            }

            saving.value = true;

            console.log("Directory:");
            console.log(selectedDirectory.value);

            if (
                window.confirm(
                    "Change storage location to " + selectedDirectory.value,
                )
            ) {
                console.log("Changed directory");
            } else {
                console.log("Canceling change");
                selectedDirectory.value = "";
                changeStoreDir.value = false;
                saving.value = false;
                return;
            }

            changeStorageLocation(selectedDirectory.value).then(() => {
                selectedDirectory.value = "";
                changeStoreDir.value = false;
                saving.value = false;
            });
        };

        // A rule that returns true if match, or an error string if not
        const directoryPathRule = (value: string) => {
            console.log("Checking rule for value: " + value);
            if (!value) return true; // If field isn't required, allow empty
            let result = windowsDirectoryRegex.test(value)
                ? true
                : "Invalid Windows directory path format";
            console.log("Passes rule: " + result);
            return result;
        };

        return {
            debts,
            useActual,
            selectedDirectory,
            saving,
            changeStoreDir,
            directoryPathRule,
            handleDirSelection,
        };
    },
});
</script>
