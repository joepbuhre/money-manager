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

            <!-- <Column selectionMode="multiple" headerStyle="width: 3rem"></Column> -->

            <Column style="width: 20%" field="code" header="Code">
                <template #editor="{ data, field }">
                    <InputText size="small" v-model="data[field]" /> </template
            ></Column>
            <Column style="width: 20%" field="name" header="Name">
                <template #editor="{ data, field }">
                    <InputText size="small" v-model="data[field]" /> </template
            ></Column>
            <Column style="width: 20%" field="category" header="Category">
                <template #editor="{ data, field }">
                    <InputText size="small" v-model="data[field]" /> </template
            ></Column>
            <Column style="width: 10%" eld="quantity" header="Quantity">
                <template #editor="{ data, field }">
                    <InputNumber
                        size="small"
                        v-model="data[field]"
                        mode="currency"
                        currency="USD"
                        locale="en-US"
                    />
                </template>
            </Column>
            <Column
                style="width: 10%"
                :rowEditor="true"
                bodyStyle="text-align:right"
            ></Column>
            <Column style="width: 3%" bodyStyle="text-align:center">
                <template #body="{ data }">
                    <button @click="removeRecord(data.id)">
                        <Trash2Icon :size="16" />
                    </button>
                </template>
            </Column>
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

const editingRows = ref<any[]>([]);
const selectedRecords = ref<any[]>([]);

const products = ref([
    {
        id: 1,
        code: "test",
        name: "test name1",
        category: "test cat",
        quantity: 10,
    },
    {
        id: 2,
        code: "test",
        name: "test name2",
        category: "test cat",
        quantity: 13,
    },
    {
        id: 3,
        code: "test",
        name: "test name3",
        category: "test cat",
        quantity: 20,
    },
]);

const removeRecord = (value: number, key: any = "id") => {
    products.value = products.value.filter((el: any) => {
        return el[key] !== value;
    });
};

const addRecord = () => {
    let i = products.value.push({
        id: -1,
        code: "",
        name: "",
        category: "",
        quantity: 0,
    });
    editingRows.value = [products.value[i - 1]];
};

const handleRowEditSave = (e: DataTableRowEditSaveEvent) => {
    products.value[e.index] = e.newData;
};
</script>
