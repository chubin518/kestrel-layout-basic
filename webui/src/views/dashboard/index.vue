<template>
  <a-spin :loading="loading" style="width: 100%">
    <div> hello golang </div>
    <ul>
      <li v-for="item in renderList" :key="item.key">
        {{ item.key }}
      </li>
    </ul>
  </a-spin>
</template>

<script lang="ts" setup>
  import useLoading from '@/hooks/loading';
  import { queryPopularList } from '@/api/dashboard';
  import type { TableData } from '@arco-design/web-vue/es/table/interface';
  import { ref } from 'vue';

  const { loading, setLoading } = useLoading();
  const renderList = ref<TableData[]>();
  const fetchData = async (contentType: string) => {
    try {
      setLoading(true);
      const { data } = await queryPopularList({ type: contentType });
      renderList.value = data;
    } catch (err) {
      // you can report use errorHandler or other
      console.error(err);
    } finally {
      setLoading(false);
    }
  };
  fetchData('text');
</script>
