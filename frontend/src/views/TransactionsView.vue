<template>
    <div class="">
        <DataTable
            editMode="row"
            dataKey="id"
            v-model:selection="selectedRecords"
            v-model:editingRows="editingRows"
            scrollable
            size="small"
            :value="products"
            :pt="{
                table: { style: 'min-width: 50rem' },
                column: {},
            }"
            @row-edit-save="handleRowEditSave"
        >
            <template #header>
                <div class="flex justify-end">
                    <PrimeButton size="small" text @click="addRecord"
                        ><PlusIcon :size="20" /> Add record</PrimeButton
                    >
                </div>
            </template>
            <!-- <Column -->

            <!-- <Column selectionMode="multiple" headerStyle="width: 3rem"></Column> -->
        </DataTable>
    </div>
</template>
<script setup lang="ts">
import { PlusIcon, Trash2Icon } from "lucide-vue-next";
import PrimeButton from "primevue/button";
import Column from "primevue/column";
import DataTable, { DataTableRowEditSaveEvent } from "primevue/datatable";
import InputNumber from "primevue/inputnumber";
import InputText from "primevue/inputtext";
import { ref } from "vue";
import { Transaction } from "../types/go/database";

const editingRows = ref<any[]>([]);
const selectedRecords = ref<any[]>([]);

const products = ref<Transaction[]>([]);

const removeRecord = (value: number, key: any = "id") => {
    products.value = products.value.filter((el: any) => {
        return el[key] !== value;
    });
};

const addRecord = () => {
    // let i = products.value.push();
    // editingRows.value = [products.value[i - 1]];
};

const handleRowEditSave = (e: DataTableRowEditSaveEvent) => {
    products.value[e.index] = e.newData;
};
</script>
