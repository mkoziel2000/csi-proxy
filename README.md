# CSI Proxy

CSI Proxy is a binary that exposes a set of gRPC APIs around storage operations
over named pipes in Windows. A container, such as CSI node plugins, can mount
the named pipes depending on operations it wants to exercise on the host and
invoke the APIs.

Each named pipe will support a specific version of an API (e.g. v1alpha1, v2beta1)
that targets a specific area of storage (e.g. disk, volume, file, SMB, iSCSI).

## Usage in a Kubernetes CSI Node Plugin DaemonSet

Deploy and start csiproxy.exe on all Windows hosts in the cluster. Next, the named
pipes can be mounted in a CSI node plugin DaemonSet YAML in the following manner:

```
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-storage-node-win
spec:
  selector:
    matchLabels:
      app: csi-driver-win
  template:
    metadata:
      labels:
        app: csi-driver-win
    spec:
      serviceAccountName: csi-node-sa
      tolerations:
      - key: "node.kubernetes.io/os"
        operator: "Equal"
        value: "win1809"
        effect: "NoSchedule"
      nodeSelector:
        kubernetes.io/os: windows
      containers:
        - name: csi-driver-registrar
          image: gke.gcr.io/csi-node-driver-registrar:win-v1
          args:
            - "--v=5"
            - "--csi-address=unix://C:\\csi\\csi.sock"
            - "--kubelet-registration-path=C:\\var\\lib\\kubelet\\plugins\\pd.csi.storage.gke.io\\csi.sock"
          env:
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: plugin-dir
              mountPath: C:\csi
            - name: registration-dir
              mountPath: C:\registration
        - name: csi-driver
          image: org/csi-driver:win-v1
          args:
            - "--v=5"
            - "--endpoint=unix:/csi/csi.sock"
          volumeMounts:
            - name: kubelet-dir
              mountPath: C:\var\lib\kubelet
            - name: plugin-dir
              mountPath: C:\csi
            - name: csi-proxy-disk-pipe
              mountPath: \\.\pipe\csi-proxy-disk-v1alpha1
            - name: csi-proxy-volume-pipe
              mountPath: \\.\pipe\csi-proxy-volume-v1alpha1
            - name: csi-proxy-filesystem-pipe
              mountPath: \\.\pipe\csi-proxy-filesystem-v1alpha1
      volumes:
        - name: csi-proxy-disk-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-disk-v1alpha1
            type: ""
        - name: csi-proxy-volume-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-volume-v1alpha1
            type: ""
        - name: csi-proxy-filesystem-pipe
          hostPath:
            path: \\.\pipe\csi-proxy-filesystem-v1alpha1
            type: ""
        - name: registration-dir
          hostPath:
            path: C:\var\lib\kubelet\plugins_registry\
            type: Directory
        - name: kubelet-dir
          hostPath:
            path: C:\var\lib\kubelet\
            type: Directory
        - name: plugin-dir
          hostPath:
            path: C:\var\lib\kubelet\plugins\csi.org.io\
            type: DirectoryOrCreate
```

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack channel](https://kubernetes.slack.com/messages/csi-windows)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-storage)

## Supported CSI Drivers

- [Azure Disk CSI Driver](https://github.com/kubernetes-sigs/azuredisk-csi-driver/tree/master/deploy/example/windows)

- [Azure File CSI Driver](https://github.com/kubernetes-sigs/azurefile-csi-driver/tree/master/deploy/example/windows)

- [Google Compute Engine Persistent Disk CSI Driver](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver)


### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).
