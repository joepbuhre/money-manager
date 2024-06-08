<template>
    <DataTable
        editMode="row"
        dataKey="id"
        v-model:editingRows="editingRows"
        scrollable
        size="small"
        :value="budgets"
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
        <Column field="created_at" header="Created At"> </Column>
        <Column field="updated_at" header="Updated At"> </Column>
        <Column field="name" header="Name">
            <template #editor="{ data, field }">
                <InputText size="small" v-model="data[field]" />
            </template>
        </Column>
        <Column field="description" header="Description">
            <template #editor="{ data, field }">
                <InputText size="small" v-model="data[field]" />
            </template>
        </Column>
        <Column field="budget_amount" header="Budget Amount">
            <template #editor="{ data, field }">
                <InputText size="small" v-model="data[field]" />
            </template>
        </Column>
        <Column field="budget_type" header="Budget Type">
            <template #body="{ data }">
                {{ (budgetTypes.find((el) => el.value === (<Budget>data).budget_type))?.name }}
            </template>
            <template #editor="{ data, field }">
                <Dropdown
                    v-model="data[field]"
                    size="small"
                    :options="budgetTypes"
                    optionLabel="name"
                    optionValue="value"
                    placeholder="Select a City"
                    class="md:w-14rem w-full"
                />
            </template>
        </Column>
        <Column
            style="width: 10%"
            :rowEditor="true"
            bodyStyle="text-align:right"
        ></Column>
    </DataTable>
</template>
<script setup lang="ts">
import { Budget } from "@/types/go/database";
import { PlusIcon } from "lucide-vue-next";
import PrimeButton from "primevue/button";
import DataTable, { DataTableRowEditSaveEvent } from "primevue/datatable";
import { ref } from "vue";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";

type T = Budget;

const budgets = ref<T[]>([]);
const editingRows = ref<(typeof budgets.value)[]>([]);

const budgetTypes = [{ name: "Normal", value: 1 }];

const removeRecord = (value: number, key: any = "id") => {
    budgets.value = budgets.value.filter((el: any) => {
        return el[key] !== value;
    });
};

const addRecord = () => {
    let i = budgets.value.push({
        id: -1,
        name: "",
        description: "",
        budget_amount: 0,
        budget_type: 1,
        created_at: null,
        updated_at: null,
    });
    editingRows.value = [budgets.value[i - 1]];
};

const handleRowEditSave = (e: DataTableRowEditSaveEvent) => {
    budgets.value[e.index] = e.newData;
};
</script>
