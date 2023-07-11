<template>
  <div>
    <h2>主机信息</h2>
    <div style="text-align: center">
      <table class="center-table">
        <tr>
          <th>主机名</th>
          <th>运行时间</th>
          <th>启动时间</th>
          <th>进程数</th>
          <th>操作系统</th>
          <th>平台</th>
          <th>平台系列</th>
          <th>平台版本</th>
          <th>内核版本</th>
          <th>内核架构</th>
          <th>虚拟化系统</th>
          <th>虚拟化角色</th>
          <th>主机ID</th>
        </tr>
        <tr>
          <td>{{ host.hostname }}</td>
          <td>{{ formatUptime(host.uptime) }}</td>
          <td>{{ formatTimestamp(host.bootTime) }}</td>
          <td>{{ host.procs }}</td>
          <td>{{ host.os }}</td>
          <td>{{ host.platform }}</td>
          <td>{{ host.platformFamily }}</td>
          <td>{{ host.platformVersion }}</td>
          <td>{{ host.kernelVersion }}</td>
          <td>{{ host.kernelArch }}</td>
          <td>{{ host.virtualizationSystem }}</td>
          <td>{{ host.virtualizationRole }}</td>
          <td>{{ host.hostid }}</td>
        </tr>
      </table>
    </div>
    <h2>内存信息</h2>
    <div style="text-align: center">
      <table class="center-table">
        <tr>
          <th>总内存</th>
          <th>可用内存</th>
          <th>已用内存</th>
          <th>已用百分比</th>
          <th>空闲内存</th>
          <th>活跃内存</th>
          <th>非活跃内存</th>
          <th>系统保留内存</th>
          <th>内存脏数据页</th>
          <th>文件系统缓存</th>
          <th>内核缓存</th>
          <th>写回缓存</th>
          <th>脏页</th>
          <th>临时写回缓存</th>
          <th>共享内存</th>
          <th>SLAB</th>
          <th>可回收SLAB</th>
          <th>不可回收SLAB</th>
          <th>页面表</th>
          <th>交换缓存</th>
        </tr>
        <tr>
          <td>{{ formatMemorySize(mem.total) }}</td>
          <td>{{ formatMemorySize(mem.available) }}</td>
          <td>{{ formatMemorySize(mem.used) }}</td>
          <td>{{ formatPercent(mem.usedPercent) }}%</td>
          <td>{{ formatMemorySize(mem.free) }}</td>
          <td>{{ formatMemorySize(mem.active) }}</td>
          <td>{{ formatMemorySize(mem.inactive) }}</td>
          <td>{{ formatMemorySize(mem.wired) }}</td>
          <td>{{ formatMemorySize(mem.laundry) }}</td>
          <td>{{ formatMemorySize(mem.buffers) }}</td>
          <td>{{ formatMemorySize(mem.cached) }}</td>
          <td>{{ formatMemorySize(mem.writeback) }}</td>
          <td>{{ formatMemorySize(mem.dirty) }}</td>
          <td>{{ formatMemorySize(mem.writebacktmp) }}</td>
          <td>{{ formatMemorySize(mem.shared) }}</td>
          <td>{{ formatMemorySize(mem.slab) }}</td>
          <td>{{ formatMemorySize(mem.sreclaimable) }}</td>
          <td>{{ formatMemorySize(mem.sunreclaim) }}</td>
          <td>{{ formatMemorySize(mem.pagetables) }}</td>
          <td>{{ formatMemorySize(mem.swapcached) }}</td>
        </tr>
      </table>
    </div>
    <h2>CPU信息</h2>
    <template v-if="cpu.length === 1">
      <div style="text-align: center">
        <table class="center-table">
          <tr v-for="(value, key) in cpu[0]" :key="key">
            <td>{{ key }}</td>
            <td>{{ value }}</td>
          </tr>
        </table>
      </div>
    </template>
    <template v-else>
      <el-table :data="cpu" stripe>
        <el-table-column prop="cpu" label="CPU"></el-table-column>
        <el-table-column prop="vendorId" label="Vendor ID"></el-table-column>
        <el-table-column prop="family" label="Family"></el-table-column>
        <el-table-column prop="model" label="Model"></el-table-column>
        <el-table-column prop="stepping" label="Stepping"></el-table-column>
        <el-table-column prop="physicalId" label="Physical ID"></el-table-column>
        <el-table-column prop="coreId" label="Core ID"></el-table-column>
        <el-table-column prop="cores" label="Cores"></el-table-column>
        <el-table-column prop="modelName" label="Model Name"></el-table-column>
        <el-table-column prop="mhz" label="MHz"></el-table-column>
        <el-table-column prop="cacheSize" label="Cache Size"></el-table-column>
        <el-table-column
          prop="flags"
          label="Flags"
          show-overflow-tooltip></el-table-column>
        <el-table-column prop="microcode" label="Microcode"></el-table-column>
      </el-table>
    </template>

    <h2>磁盘信息</h2>
    <div style="text-align: center">
      <table class="center-table">
        <tr>
          <th>路径</th>
          <th>文件系统类型</th>
          <th>总空间</th>
          <th>可用空间</th>
          <th>已用空间</th>
          <th>已用百分比</th>
          <th>索引节点总数</th>
          <th>已用索引节点数</th>
          <th>可用索引节点数</th>
          <th>已用索引节点百分比</th>
        </tr>
        <tr>
          <td>{{ disk.path }}</td>
          <td>{{ disk.fstype }}</td>
          <td>{{ formatMemorySize(disk.total) }}</td>
          <td>{{ formatMemorySize(disk.free) }}</td>
          <td>{{ formatMemorySize(disk.used) }}</td>
          <td>{{ formatPercent(disk.usedPercent) }}%</td>
          <td>{{ disk.inodesTotal }}</td>
          <td>{{ disk.inodesUsed }}</td>
          <td>{{ disk.inodesFree }}</td>
          <td>{{ formatPercent(disk.inodesUsedPercent) }}%</td>
        </tr>
      </table>
    </div>

    <h2>系统负载</h2>
    <div style="text-align: center">
      <table class="center-table">
        <tr>
          <th>1分钟平均负载</th>
          <th>5分钟平均负载</th>
          <th>15分钟平均负载</th>
        </tr>
        <tr>
          <td>{{ formatPercent(load.load1) }}%</td>
          <td>{{ formatPercent(load.load5) }}%</td>
          <td>{{ formatPercent(load.load15) }}%</td>
        </tr>
      </table>
    </div>

    <h2>网络信息</h2>

    <div style="text-align: center; width: '70%'">
      <el-table :data="net">
        <el-table-column prop="name" label="接口名称"></el-table-column>
        <el-table-column prop="bytesSent" label="发送字节数"></el-table-column>
        <el-table-column prop="bytesRecv" label="接收字节"></el-table-column>
        <el-table-column prop="packetsSent" label="发送数据包数"></el-table-column>
        <el-table-column prop="packetsRecv" label="接收数据包数"></el-table-column>
        <el-table-column prop="errin" label="输入错误数"></el-table-column>
        <el-table-column prop="errout" label="输出错误数"></el-table-column>
        <el-table-column prop="dropin" label="输入丢弃数"></el-table-column>
        <el-table-column prop="dropout" label="输出丢弃数"></el-table-column>
        <el-table-column prop="fifoin" label="输入FIFO错误数"></el-table-column>
        <el-table-column prop="fifoout" label="输出FIFO错误数"></el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      cpu: [],
      disk: {},
      host: {},
      load: {},
      mem: {},
      net: [],
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    formatUptime(uptime) {
      const days = Math.floor(uptime / 86400)
      const hours = Math.floor((uptime % 86400) / 3600)
      const minutes = Math.floor((uptime % 3600) / 60)
      const seconds = uptime % 60

      return `${days}天 ${hours}小时 ${minutes}分钟 ${seconds}秒`
    },
    formatTimestamp(timestamp) {
      const date = new Date(timestamp * 1000)
      const formattedDate = date.toLocaleString()
      return formattedDate
    },
    formatPercent(percen) {
      return percen.toFixed(2)
    },
    formatMemorySize(size) {
      const KB = 1024
      const MB = KB * 1024
      const GB = MB * 1024

      if (size >= GB) {
        return `${(size / GB).toFixed(2)} GB`
      } else if (size >= MB) {
        return `${(size / MB).toFixed(2)} MB`
      } else {
        return `${(size / KB).toFixed(2)} KB`
      }
    },
    fetchData() {
      axios
        .get('/info')
        .then((response) => {
          const data = response.data
          this.cpu = data.cpu
          this.disk = data.disk
          this.host = data.host
          this.load = data.load
          this.mem = data.mem
          this.net = data.net
        })
        .catch((error) => {
          console.error('Error fetching data:', error)
        })
    },
  },
}
</script>

<style scoped>
.center-table {
  margin: 0 auto;
  width: 70%;
  border-collapse: collapse; /* 合并边框 */
}

.center-table th,
.center-table td {
  border: 1px solid #ddd; /* 单元格边框 */
  padding: 8px; /* 单元格内边距 */
}

.center-table th {
  background-color: #f2f2f2; /* 表头背景色 */
}

.center-table tr:nth-child(even) {
  background-color: #f2f2f2; /* 偶数行的背景色 */
}
</style>
