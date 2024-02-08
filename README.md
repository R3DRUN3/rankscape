# CNCF Landscape Ranking 🖼 📈
<img src="images/lndscp.jpg" alt="cncf" height="150" width="150">  


<br/>

Ranking of *Cloud Native Computing Foundation* landscape projects via *Criticality Score*.  


## Abstract
The [CNCF projects](https://github.com/cncf/landscape) are growing in number, making it increasingly difficult to keep track of them all.  
While some aim to address similar challenges, others operate in distinct domains, yet they all share a common attribute:  
**critical importance to modern infrastructures**.  
<br/>

This repository is driven by the primary rationale of maintaining an up-to-date ranking of both *incubating* and *graduated* projects, leveraging the [**criticality score**](https://github.com/ossf/criticality_score) as a benchmark.  

## How It Works
This repository contains a *Go* script that performs the following tasks:
1. Retrieves the [official list of projects](https://github.com/cncf/landscape/blob/master/landscape.yml) from the *CNCF landscape*.
1. For each project in the `incubating` and/or `graduated` state, calculates its score using [this](https://github.com/engelsjk/criticalityscore/) port of the *criticality score*.
1. Sorts the projects by score from highest to lowest.
1. Updates this same *README* by replacing the placeholder `<!--TABLE_PLACEHOLDER-->` with the HTML table of the final ranking (see it below).  


> [!NOTE]  
> Each project is processed sequentially; hence, script execution takes a few minutes to complete.  
> I attempted parallelization using Go routines, but I was immediately rate-limited by the GitHub APIs.  


<br/>


## Ranking
<table><tr><th>Rank</th><th>Name</th><th>Score</th><th>Scored On</th></tr><tr><td>1</td><td><a href="https://github.com/kubernetes/kubernetes">kubernetes</a></td><td>0.703920</td><td>Wed Feb  7 19:25:51 2024</td></tr><tr><td>2</td><td><a href="https://github.com/backstage/backstage">backstage</a></td><td>0.699100</td><td>Wed Feb  7 19:27:44 2024</td></tr><tr><td>3</td><td><a href="https://github.com/envoyproxy/envoy">envoy</a></td><td>0.685070</td><td>Wed Feb  7 19:26:40 2024</td></tr><tr><td>4</td><td><a href="https://github.com/cilium/cilium">cilium</a></td><td>0.678210</td><td>Wed Feb  7 19:25:19 2024</td></tr><tr><td>5</td><td><a href="https://github.com/istio/istio">istio</a></td><td>0.675740</td><td>Wed Feb  7 19:26:50 2024</td></tr><tr><td>6</td><td><a href="https://github.com/prometheus/prometheus">prometheus</a></td><td>0.668490</td><td>Wed Feb  7 19:29:16 2024</td></tr><tr><td>7</td><td><a href="https://github.com/keycloak/keycloak">keycloak</a></td><td>0.666890</td><td>Wed Feb  7 19:23:41 2024</td></tr><tr><td>8</td><td><a href="https://github.com/containerd/containerd">containerd</a></td><td>0.650040</td><td>Wed Feb  7 19:24:45 2024</td></tr><tr><td>9</td><td><a href="https://github.com/rook/rook">rook</a></td><td>0.649060</td><td>Wed Feb  7 19:24:40 2024</td></tr><tr><td>10</td><td><a href="https://github.com/etcd-io/etcd">etcd</a></td><td>0.644870</td><td>Wed Feb  7 19:26:29 2024</td></tr><tr><td>11</td><td><a href="https://github.com/grpc/grpc">grpc</a></td><td>0.628070</td><td>Wed Feb  7 19:26:33 2024</td></tr><tr><td>12</td><td><a href="https://github.com/kubevirt/kubevirt">kubevirt</a></td><td>0.618730</td><td>Wed Feb  7 19:28:07 2024</td></tr><tr><td>13</td><td><a href="https://github.com/helm/helm">helm</a></td><td>0.615450</td><td>Wed Feb  7 19:27:59 2024</td></tr><tr><td>14</td><td><a href="https://github.com/knative/serving">serving</a></td><td>0.612150</td><td>Wed Feb  7 19:25:37 2024</td></tr><tr><td>15</td><td><a href="https://github.com/kubeedge/kubeedge">kubeedge</a></td><td>0.609840</td><td>Wed Feb  7 19:23:19 2024</td></tr><tr><td>16</td><td><a href="https://github.com/vitessio/vitess">vitess</a></td><td>0.605770</td><td>Wed Feb  7 19:27:02 2024</td></tr><tr><td>17</td><td><a href="https://github.com/dapr/dapr">dapr</a></td><td>0.603690</td><td>Wed Feb  7 19:28:58 2024</td></tr><tr><td>18</td><td><a href="https://github.com/open-policy-agent/opa">opa</a></td><td>0.595130</td><td>Wed Feb  7 19:24:19 2024</td></tr><tr><td>19</td><td><a href="https://github.com/argoproj/argo-cd">argo-cd</a></td><td>0.587120</td><td>Wed Feb  7 19:28:14 2024</td></tr><tr><td>20</td><td><a href="https://github.com/kedacore/keda">keda</a></td><td>0.584200</td><td>Wed Feb  7 19:29:01 2024</td></tr><tr><td>21</td><td><a href="https://github.com/cert-manager/cert-manager">cert-manager</a></td><td>0.581900</td><td>Wed Feb  7 19:23:30 2024</td></tr><tr><td>22</td><td><a href="https://github.com/goharbor/harbor">harbor</a></td><td>0.572720</td><td>Wed Feb  7 19:23:27 2024</td></tr><tr><td>23</td><td><a href="https://github.com/thanos-io/thanos">thanos</a></td><td>0.568450</td><td>Wed Feb  7 19:29:20 2024</td></tr><tr><td>24</td><td><a href="https://github.com/knative/serving">serving</a></td><td>0.566620</td><td>Wed Feb  7 19:29:05 2024</td></tr><tr><td>25</td><td><a href="https://github.com/cloud-custodian/cloud-custodian">cloud-custodian</a></td><td>0.566550</td><td>Wed Feb  7 19:23:16 2024</td></tr><tr><td>26</td><td><a href="https://github.com/dapr/dapr">dapr</a></td><td>0.558390</td><td>Wed Feb  7 19:27:55 2024</td></tr><tr><td>27</td><td><a href="https://github.com/projectcontour/contour">contour</a></td><td>0.556520</td><td>Wed Feb  7 19:26:36 2024</td></tr><tr><td>28</td><td><a href="https://github.com/tikv/tikv">tikv</a></td><td>0.554170</td><td>Wed Feb  7 19:26:58 2024</td></tr><tr><td>29</td><td><a href="https://github.com/kedacore/keda">keda</a></td><td>0.547960</td><td>Wed Feb  7 19:25:33 2024</td></tr><tr><td>30</td><td><a href="https://github.com/karmada-io/karmada">karmada</a></td><td>0.546660</td><td>Wed Feb  7 19:25:29 2024</td></tr><tr><td>31</td><td><a href="https://github.com/emissary-ingress/emissary">emissary</a></td><td>0.543350</td><td>Wed Feb  7 19:26:44 2024</td></tr><tr><td>32</td><td><a href="https://github.com/cri-o/cri-o">cri-o</a></td><td>0.540460</td><td>Wed Feb  7 19:24:48 2024</td></tr><tr><td>33</td><td><a href="https://github.com/longhorn/longhorn">longhorn</a></td><td>0.539230</td><td>Wed Feb  7 19:24:37 2024</td></tr><tr><td>34</td><td><a href="https://github.com/open-telemetry/opentelemetry-java">opentelemetry-java</a></td><td>0.539160</td><td>Wed Feb  7 19:30:02 2024</td></tr><tr><td>35</td><td><a href="https://github.com/kyverno/kyverno">kyverno</a></td><td>0.533690</td><td>Wed Feb  7 19:23:44 2024</td></tr><tr><td>36</td><td><a href="https://github.com/nats-io/nats-server">nats-server</a></td><td>0.533160</td><td>Wed Feb  7 19:27:36 2024</td></tr><tr><td>37</td><td><a href="https://github.com/falcosecurity/falco">falco</a></td><td>0.531280</td><td>Wed Feb  7 19:23:34 2024</td></tr><tr><td>38</td><td><a href="https://github.com/cortexproject/cortex">cortex</a></td><td>0.530290</td><td>Wed Feb  7 19:29:08 2024</td></tr><tr><td>39</td><td><a href="https://github.com/linkerd/linkerd2">linkerd2</a></td><td>0.525180</td><td>Wed Feb  7 19:26:54 2024</td></tr><tr><td>40</td><td><a href="https://github.com/kubeflow/kubeflow">kubeflow</a></td><td>0.523000</td><td>Wed Feb  7 19:25:40 2024</td></tr><tr><td>41</td><td><a href="https://github.com/jaegertracing/jaeger">jaeger</a></td><td>0.522200</td><td>Wed Feb  7 19:29:58 2024</td></tr><tr><td>42</td><td><a href="https://github.com/strimzi/strimzi-kafka-operator">strimzi-kafka-operator</a></td><td>0.521910</td><td>Wed Feb  7 19:27:40 2024</td></tr><tr><td>43</td><td><a href="https://github.com/volcano-sh/volcano">volcano</a></td><td>0.511580</td><td>Wed Feb  7 19:25:55 2024</td></tr><tr><td>44</td><td><a href="https://github.com/litmuschaos/litmus">litmus</a></td><td>0.508760</td><td>Wed Feb  7 19:30:14 2024</td></tr><tr><td>45</td><td><a href="https://github.com/crossplane/crossplane">crossplane</a></td><td>0.508070</td><td>Wed Feb  7 19:25:26 2024</td></tr><tr><td>46</td><td><a href="https://github.com/operator-framework/operator-sdk">operator-sdk</a></td><td>0.502880</td><td>Wed Feb  7 19:28:10 2024</td></tr><tr><td>47</td><td><a href="https://github.com/fluxcd/flux2">flux2</a></td><td>0.500920</td><td>Wed Feb  7 19:28:45 2024</td></tr><tr><td>48</td><td><a href="https://github.com/buildpacks/pack">pack</a></td><td>0.500810</td><td>Wed Feb  7 19:27:47 2024</td></tr><tr><td>49</td><td><a href="https://github.com/fluent/fluentd">fluentd</a></td><td>0.496930</td><td>Wed Feb  7 19:29:54 2024</td></tr><tr><td>50</td><td><a href="https://github.com/chaos-mesh/chaos-mesh">chaos-mesh</a></td><td>0.491710</td><td>Wed Feb  7 19:30:10 2024</td></tr><tr><td>51</td><td><a href="https://github.com/spiffe/spire">spire</a></td><td>0.490630</td><td>Wed Feb  7 19:24:30 2024</td></tr><tr><td>52</td><td><a href="https://github.com/kubevela/kubevela">kubevela</a></td><td>0.490470</td><td>Wed Feb  7 19:28:03 2024</td></tr><tr><td>53</td><td><a href="https://github.com/coredns/coredns">coredns</a></td><td>0.486780</td><td>Wed Feb  7 19:25:59 2024</td></tr><tr><td>54</td><td><a href="https://github.com/dragonflyoss/Dragonfly2">Dragonfly2</a></td><td>0.479050</td><td>Wed Feb  7 19:23:23 2024</td></tr><tr><td>55</td><td><a href="https://github.com/openkruise/kruise">kruise</a></td><td>0.476520</td><td>Wed Feb  7 19:28:52 2024</td></tr><tr><td>56</td><td><a href="https://github.com/containernetworking/cni">cni</a></td><td>0.471110</td><td>Wed Feb  7 19:25:23 2024</td></tr><tr><td>57</td><td><a href="https://github.com/cloudevents/spec">spec</a></td><td>0.467230</td><td>Wed Feb  7 19:27:05 2024</td></tr><tr><td>58</td><td><a href="https://github.com/keptn/lifecycle-toolkit">lifecycle-toolkit</a></td><td>0.465110</td><td>Wed Feb  7 19:28:48 2024</td></tr><tr><td>59</td><td><a href="https://github.com/theupdateframework/python-tuf">python-tuf</a></td><td>0.463820</td><td>Wed Feb  7 19:24:22 2024</td></tr><tr><td>60</td><td><a href="https://github.com/in-toto/in-toto">in-toto</a></td><td>0.436140</td><td>Wed Feb  7 19:23:37 2024</td></tr><tr><td>61</td><td><a href="https://github.com/cubeFS/cubefs">cubefs</a></td><td>0.412790</td><td>Wed Feb  7 19:24:33 2024</td></tr><tr><td>62</td><td><a href="https://github.com/notaryproject/notation">notation</a></td><td>0.403930</td><td>Wed Feb  7 19:24:15 2024</td></tr><tr><td>63</td><td><a href="https://github.com/open-feature/spec">spec</a></td><td>0.395100</td><td>Wed Feb  7 19:29:51 2024</td></tr><tr><td>64</td><td><a href="https://github.com/spiffe/spiffe">spiffe</a></td><td>0.352460</td><td>Wed Feb  7 19:24:26 2024</td></tr><tr><td>65</td><td><a href="https://github.com/opentracing/opentracing-go">opentracing-go</a></td><td>0.259350</td><td>Wed Feb  7 19:30:06 2024</td></tr><tr><td>66</td><td><a href="https://github.com/OpenObservability/OpenMetrics">OpenMetrics</a></td><td>0.250390</td><td>Wed Feb  7 19:29:11 2024</td></tr></table>