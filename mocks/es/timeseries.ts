import { createPromiseClient, createRouterTransport } from '@connectrpc/connect'
import { TimeseriesService } from '../../build/es/timeseries/timeseries_connect'
import {
  DataPoint, DeviceMetric, MetricInfo, FlushResponse, Metric, ReadResponse, StatResponse, WriteBulkResponse, WriteResponse
} from '../../build/es/timeseries/timeseries_pb'
import { DevicesService } from '../../build/es/node/node_connect'
import { QueryRequest } from '../../build/es/node/node_pb'
import { transport as devicesTransport } from './devices.ts'

export const transport = createRouterTransport(({ service }) => {
  const devicesApi = createPromiseClient(DevicesService, devicesTransport)
  const metrics = new Map<string, Map<string, DataPoint[]>>()

  function getField() {
    return [100, 90, 80, 70, 60, 50, 40, 30, 20, 10].map((num) => {
      const ts = BigInt(Date.now() - num * 3600 * 1000)
      const value = Math.floor(Math.random() * 101)

      return new DataPoint({ ts, value })
    })
  }

  function getMetrics(uuid: string) {
    const metricsList: Metric[] = []

    metrics.get(uuid)?.forEach((value, name) => {
      const dataPoints = value.length
      let earliest = value[0]
      let latest = value[0]

      value.forEach((point) => {
        if (point.ts < earliest.ts) {
          earliest = point
        }

        if (point.ts > latest.ts) {
          latest = point
        }
      })

      metricsList.push(new Metric({ dataPoints, name, earliest, latest }))
    })

    return metricsList
  }

  devicesApi.list(new QueryRequest()).then(({ devices }) => {
    devices.forEach(({ uuid }) => {
      metrics.set(uuid, new Map([
        ['fuel', getField()], ['temp', getField()]
      ]))
    })
  })

  service(TimeseriesService, {
    read(request) {
      const fieldsInfo: MetricInfo[] = []

      if (!metrics.has(request.device)) {
        metrics.set(request.device, new Map())
      }

      request.metrics?.forEach((metric) => {
        const fields = metrics.get(request.device)

        const dataPoints = fields?.get(metric)?.filter(({ ts }) =>
          ts >= request.from && (request.to) ? ts <= request.to : true
        )

        fieldsInfo.push(new MetricInfo({ metric, dataPoints }))
      })

      return new ReadResponse({ metricsInfo: fieldsInfo })
    },
    async stat(request) {
      await new Promise((resolve) => setTimeout(resolve, 300))
      const deviceMetrics: DeviceMetric[] = []

      if (request.namespace) {
        const { devices } = await devicesApi.list({ namespace: request.namespace })

        devices.forEach(({ uuid }) => {
          deviceMetrics.push(new DeviceMetric({
            device: uuid, metrics: getMetrics(uuid)
          }))
        })
      } else if (request.device) {
        deviceMetrics.push(new DeviceMetric({
          device: request.device, metrics: getMetrics(request.device)
        }))
      } else {
        for (const device of metrics.keys()) {
          deviceMetrics.push(new DeviceMetric({
            device, metrics: getMetrics(device)
          }))
        }
      }

      return new StatResponse({ deviceMetrics })
    },
    write(request) {
      const dataPoint = request.dataPoint ?? new DataPoint()
      const metric = new Map([[request.metric, [dataPoint]]])
      const value = metrics.get(request.device)

      if (!value) {
        metrics.set(request.device, metric)
      } else if (!value.has(request.metric)) {
        value.set(request.metric, [dataPoint])
      } else {
        value.get(request.metric)?.push(dataPoint)
      }

      return new WriteResponse({ result: true, ts: BigInt(Date.now()) })
    },
    writeBulk(request) {
      const metric = new Map([[request.metric, request.dataPoints]])
      const value = metrics.get(request.device)

      if (!value) {
        metrics.set(request.device, metric)
      } else if (!value.has(request.metric)) {
        value.set(request.metric, request.dataPoints)
      } else {
        value.get(request.metric)?.push(...request.dataPoints)
      }

      return new WriteBulkResponse({ result: true })
    },
    async flush(request) {
      if (request.namespace) {
        const { devices } = await devicesApi.list(new QueryRequest({
          namespace: request.namespace
        }))

        devices.forEach(({ uuid }) => {
          metrics.delete(uuid)
        })
      } else if (request.device) {
        metrics.delete(request.device as string)
      } else {
        metrics.clear()
      }

      return new FlushResponse({ result: true })
    }
  })
})
