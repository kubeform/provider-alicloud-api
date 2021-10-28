/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	reflect "reflect"
	sync "sync"
	time "time"

	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	actiontrail "kubeform.dev/provider-alicloud-api/client/informers/externalversions/actiontrail"
	adb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/adb"
	alb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/alb"
	alidns "kubeform.dev/provider-alicloud-api/client/informers/externalversions/alidns"
	alikafka "kubeform.dev/provider-alicloud-api/client/informers/externalversions/alikafka"
	amqp "kubeform.dev/provider-alicloud-api/client/informers/externalversions/amqp"
	apigateway "kubeform.dev/provider-alicloud-api/client/informers/externalversions/apigateway"
	arms "kubeform.dev/provider-alicloud-api/client/informers/externalversions/arms"
	auto "kubeform.dev/provider-alicloud-api/client/informers/externalversions/auto"
	bastionhost "kubeform.dev/provider-alicloud-api/client/informers/externalversions/bastionhost"
	brain "kubeform.dev/provider-alicloud-api/client/informers/externalversions/brain"
	cas "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cas"
	cassandra "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cassandra"
	cddc "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cddc"
	cdn "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cdn"
	cen "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cen"
	click "kubeform.dev/provider-alicloud-api/client/informers/externalversions/click"
	cloud "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cloud"
	cloudauth "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cloudauth"
	cms "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cms"
	common "kubeform.dev/provider-alicloud-api/client/informers/externalversions/common"
	config "kubeform.dev/provider-alicloud-api/client/informers/externalversions/config"
	container "kubeform.dev/provider-alicloud-api/client/informers/externalversions/container"
	copy "kubeform.dev/provider-alicloud-api/client/informers/externalversions/copy"
	cr "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cr"
	cs "kubeform.dev/provider-alicloud-api/client/informers/externalversions/cs"
	data "kubeform.dev/provider-alicloud-api/client/informers/externalversions/data"
	database "kubeform.dev/provider-alicloud-api/client/informers/externalversions/database"
	datahub "kubeform.dev/provider-alicloud-api/client/informers/externalversions/datahub"
	db "kubeform.dev/provider-alicloud-api/client/informers/externalversions/db"
	dbfs "kubeform.dev/provider-alicloud-api/client/informers/externalversions/dbfs"
	dcdn "kubeform.dev/provider-alicloud-api/client/informers/externalversions/dcdn"
	ddosbgp "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ddosbgp"
	ddoscoo "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ddoscoo"
	dfs "kubeform.dev/provider-alicloud-api/client/informers/externalversions/dfs"
	direct "kubeform.dev/provider-alicloud-api/client/informers/externalversions/direct"
	disk "kubeform.dev/provider-alicloud-api/client/informers/externalversions/disk"
	dms "kubeform.dev/provider-alicloud-api/client/informers/externalversions/dms"
	dns "kubeform.dev/provider-alicloud-api/client/informers/externalversions/dns"
	drds "kubeform.dev/provider-alicloud-api/client/informers/externalversions/drds"
	dts "kubeform.dev/provider-alicloud-api/client/informers/externalversions/dts"
	eais "kubeform.dev/provider-alicloud-api/client/informers/externalversions/eais"
	ecd "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ecd"
	eci "kubeform.dev/provider-alicloud-api/client/informers/externalversions/eci"
	ecp "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ecp"
	ecs "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ecs"
	edas "kubeform.dev/provider-alicloud-api/client/informers/externalversions/edas"
	ehpc "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ehpc"
	eip "kubeform.dev/provider-alicloud-api/client/informers/externalversions/eip"
	eipanycast "kubeform.dev/provider-alicloud-api/client/informers/externalversions/eipanycast"
	elasticsearch "kubeform.dev/provider-alicloud-api/client/informers/externalversions/elasticsearch"
	emr "kubeform.dev/provider-alicloud-api/client/informers/externalversions/emr"
	ens "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ens"
	ess "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ess"
	event "kubeform.dev/provider-alicloud-api/client/informers/externalversions/event"
	express "kubeform.dev/provider-alicloud-api/client/informers/externalversions/express"
	fc "kubeform.dev/provider-alicloud-api/client/informers/externalversions/fc"
	fnf "kubeform.dev/provider-alicloud-api/client/informers/externalversions/fnf"
	forward "kubeform.dev/provider-alicloud-api/client/informers/externalversions/forward"
	ga "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ga"
	gpdb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/gpdb"
	graph "kubeform.dev/provider-alicloud-api/client/informers/externalversions/graph"
	havip "kubeform.dev/provider-alicloud-api/client/informers/externalversions/havip"
	hbase "kubeform.dev/provider-alicloud-api/client/informers/externalversions/hbase"
	hbr "kubeform.dev/provider-alicloud-api/client/informers/externalversions/hbr"
	image "kubeform.dev/provider-alicloud-api/client/informers/externalversions/image"
	imm "kubeform.dev/provider-alicloud-api/client/informers/externalversions/imm"
	imp "kubeform.dev/provider-alicloud-api/client/informers/externalversions/imp"
	instance "kubeform.dev/provider-alicloud-api/client/informers/externalversions/instance"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	iot "kubeform.dev/provider-alicloud-api/client/informers/externalversions/iot"
	key "kubeform.dev/provider-alicloud-api/client/informers/externalversions/key"
	kms "kubeform.dev/provider-alicloud-api/client/informers/externalversions/kms"
	kvstore "kubeform.dev/provider-alicloud-api/client/informers/externalversions/kvstore"
	launch "kubeform.dev/provider-alicloud-api/client/informers/externalversions/launch"
	lindorm "kubeform.dev/provider-alicloud-api/client/informers/externalversions/lindorm"
	log "kubeform.dev/provider-alicloud-api/client/informers/externalversions/log"
	logtail "kubeform.dev/provider-alicloud-api/client/informers/externalversions/logtail"
	market "kubeform.dev/provider-alicloud-api/client/informers/externalversions/market"
	maxcompute "kubeform.dev/provider-alicloud-api/client/informers/externalversions/maxcompute"
	mhub "kubeform.dev/provider-alicloud-api/client/informers/externalversions/mhub"
	mns "kubeform.dev/provider-alicloud-api/client/informers/externalversions/mns"
	mongodb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/mongodb"
	msc "kubeform.dev/provider-alicloud-api/client/informers/externalversions/msc"
	mse "kubeform.dev/provider-alicloud-api/client/informers/externalversions/mse"
	nas "kubeform.dev/provider-alicloud-api/client/informers/externalversions/nas"
	nat "kubeform.dev/provider-alicloud-api/client/informers/externalversions/nat"
	network "kubeform.dev/provider-alicloud-api/client/informers/externalversions/network"
	ons "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ons"
	oos "kubeform.dev/provider-alicloud-api/client/informers/externalversions/oos"
	open "kubeform.dev/provider-alicloud-api/client/informers/externalversions/open"
	oss "kubeform.dev/provider-alicloud-api/client/informers/externalversions/oss"
	ots "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ots"
	polardb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/polardb"
	privatelink "kubeform.dev/provider-alicloud-api/client/informers/externalversions/privatelink"
	pvtz "kubeform.dev/provider-alicloud-api/client/informers/externalversions/pvtz"
	quick "kubeform.dev/provider-alicloud-api/client/informers/externalversions/quick"
	quotas "kubeform.dev/provider-alicloud-api/client/informers/externalversions/quotas"
	ram "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ram"
	rdc "kubeform.dev/provider-alicloud-api/client/informers/externalversions/rdc"
	rds "kubeform.dev/provider-alicloud-api/client/informers/externalversions/rds"
	reserved "kubeform.dev/provider-alicloud-api/client/informers/externalversions/reserved"
	resource "kubeform.dev/provider-alicloud-api/client/informers/externalversions/resource"
	ros "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ros"
	route "kubeform.dev/provider-alicloud-api/client/informers/externalversions/route"
	router "kubeform.dev/provider-alicloud-api/client/informers/externalversions/router"
	sae "kubeform.dev/provider-alicloud-api/client/informers/externalversions/sae"
	sag "kubeform.dev/provider-alicloud-api/client/informers/externalversions/sag"
	scdn "kubeform.dev/provider-alicloud-api/client/informers/externalversions/scdn"
	sddp "kubeform.dev/provider-alicloud-api/client/informers/externalversions/sddp"
	security "kubeform.dev/provider-alicloud-api/client/informers/externalversions/security"
	service "kubeform.dev/provider-alicloud-api/client/informers/externalversions/service"
	simple "kubeform.dev/provider-alicloud-api/client/informers/externalversions/simple"
	slb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/slb"
	snapshot "kubeform.dev/provider-alicloud-api/client/informers/externalversions/snapshot"
	snat "kubeform.dev/provider-alicloud-api/client/informers/externalversions/snat"
	ssl "kubeform.dev/provider-alicloud-api/client/informers/externalversions/ssl"
	subnet "kubeform.dev/provider-alicloud-api/client/informers/externalversions/subnet"
	tsdb "kubeform.dev/provider-alicloud-api/client/informers/externalversions/tsdb"
	video "kubeform.dev/provider-alicloud-api/client/informers/externalversions/video"
	vod "kubeform.dev/provider-alicloud-api/client/informers/externalversions/vod"
	vpc "kubeform.dev/provider-alicloud-api/client/informers/externalversions/vpc"
	vpn "kubeform.dev/provider-alicloud-api/client/informers/externalversions/vpn"
	vswitch "kubeform.dev/provider-alicloud-api/client/informers/externalversions/vswitch"
	waf "kubeform.dev/provider-alicloud-api/client/informers/externalversions/waf"
	yundun "kubeform.dev/provider-alicloud-api/client/informers/externalversions/yundun"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           versioned.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[v1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.namespace = namespace
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client versioned.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
// Deprecated: Please use NewSharedInformerFactoryWithOptions instead
func NewFilteredSharedInformerFactory(client versioned.Interface, defaultResync time.Duration, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace), WithTweakListOptions(tweakListOptions))
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client versioned.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        v1.NamespaceAll,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Actiontrail() actiontrail.Interface
	Adb() adb.Interface
	Alb() alb.Interface
	Alidns() alidns.Interface
	Alikafka() alikafka.Interface
	Amqp() amqp.Interface
	Apigateway() apigateway.Interface
	Arms() arms.Interface
	Auto() auto.Interface
	Bastionhost() bastionhost.Interface
	Brain() brain.Interface
	Cas() cas.Interface
	Cassandra() cassandra.Interface
	Cddc() cddc.Interface
	Cdn() cdn.Interface
	Cen() cen.Interface
	Click() click.Interface
	Cloud() cloud.Interface
	Cloudauth() cloudauth.Interface
	Cms() cms.Interface
	Common() common.Interface
	Config() config.Interface
	Container() container.Interface
	Copy() copy.Interface
	Cr() cr.Interface
	Cs() cs.Interface
	Data() data.Interface
	Database() database.Interface
	Datahub() datahub.Interface
	Db() db.Interface
	Dbfs() dbfs.Interface
	Dcdn() dcdn.Interface
	Ddosbgp() ddosbgp.Interface
	Ddoscoo() ddoscoo.Interface
	Dfs() dfs.Interface
	Direct() direct.Interface
	Disk() disk.Interface
	Dms() dms.Interface
	Dns() dns.Interface
	Drds() drds.Interface
	Dts() dts.Interface
	Eais() eais.Interface
	Ecd() ecd.Interface
	Eci() eci.Interface
	Ecp() ecp.Interface
	Ecs() ecs.Interface
	Edas() edas.Interface
	Ehpc() ehpc.Interface
	Eip() eip.Interface
	Eipanycast() eipanycast.Interface
	Elasticsearch() elasticsearch.Interface
	Emr() emr.Interface
	Ens() ens.Interface
	Ess() ess.Interface
	Event() event.Interface
	Express() express.Interface
	Fc() fc.Interface
	Fnf() fnf.Interface
	Forward() forward.Interface
	Ga() ga.Interface
	Gpdb() gpdb.Interface
	Graph() graph.Interface
	Havip() havip.Interface
	Hbase() hbase.Interface
	Hbr() hbr.Interface
	Image() image.Interface
	Imm() imm.Interface
	Imp() imp.Interface
	Instance() instance.Interface
	Iot() iot.Interface
	Key() key.Interface
	Kms() kms.Interface
	Kvstore() kvstore.Interface
	Launch() launch.Interface
	Lindorm() lindorm.Interface
	Log() log.Interface
	Logtail() logtail.Interface
	Market() market.Interface
	Maxcompute() maxcompute.Interface
	Mhub() mhub.Interface
	Mns() mns.Interface
	Mongodb() mongodb.Interface
	Msc() msc.Interface
	Mse() mse.Interface
	Nas() nas.Interface
	Nat() nat.Interface
	Network() network.Interface
	Ons() ons.Interface
	Oos() oos.Interface
	Open() open.Interface
	Oss() oss.Interface
	Ots() ots.Interface
	Polardb() polardb.Interface
	Privatelink() privatelink.Interface
	Pvtz() pvtz.Interface
	Quick() quick.Interface
	Quotas() quotas.Interface
	Ram() ram.Interface
	Rdc() rdc.Interface
	Rds() rds.Interface
	Reserved() reserved.Interface
	Resource() resource.Interface
	Ros() ros.Interface
	Route() route.Interface
	Router() router.Interface
	Sae() sae.Interface
	Sag() sag.Interface
	Scdn() scdn.Interface
	Sddp() sddp.Interface
	Security() security.Interface
	Service() service.Interface
	Simple() simple.Interface
	Slb() slb.Interface
	Snapshot() snapshot.Interface
	Snat() snat.Interface
	Ssl() ssl.Interface
	Subnet() subnet.Interface
	Tsdb() tsdb.Interface
	Video() video.Interface
	Vod() vod.Interface
	Vpc() vpc.Interface
	Vpn() vpn.Interface
	Vswitch() vswitch.Interface
	Waf() waf.Interface
	Yundun() yundun.Interface
}

func (f *sharedInformerFactory) Actiontrail() actiontrail.Interface {
	return actiontrail.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Adb() adb.Interface {
	return adb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Alb() alb.Interface {
	return alb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Alidns() alidns.Interface {
	return alidns.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Alikafka() alikafka.Interface {
	return alikafka.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Amqp() amqp.Interface {
	return amqp.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apigateway() apigateway.Interface {
	return apigateway.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Arms() arms.Interface {
	return arms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Auto() auto.Interface {
	return auto.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Bastionhost() bastionhost.Interface {
	return bastionhost.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Brain() brain.Interface {
	return brain.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cas() cas.Interface {
	return cas.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cassandra() cassandra.Interface {
	return cassandra.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cddc() cddc.Interface {
	return cddc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cdn() cdn.Interface {
	return cdn.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cen() cen.Interface {
	return cen.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Click() click.Interface {
	return click.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloud() cloud.Interface {
	return cloud.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudauth() cloudauth.Interface {
	return cloudauth.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cms() cms.Interface {
	return cms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Common() common.Interface {
	return common.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Config() config.Interface {
	return config.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Container() container.Interface {
	return container.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Copy() copy.Interface {
	return copy.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cr() cr.Interface {
	return cr.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cs() cs.Interface {
	return cs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Data() data.Interface {
	return data.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Database() database.Interface {
	return database.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Datahub() datahub.Interface {
	return datahub.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Db() db.Interface {
	return db.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dbfs() dbfs.Interface {
	return dbfs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dcdn() dcdn.Interface {
	return dcdn.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ddosbgp() ddosbgp.Interface {
	return ddosbgp.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ddoscoo() ddoscoo.Interface {
	return ddoscoo.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dfs() dfs.Interface {
	return dfs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Direct() direct.Interface {
	return direct.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Disk() disk.Interface {
	return disk.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dms() dms.Interface {
	return dms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dns() dns.Interface {
	return dns.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Drds() drds.Interface {
	return drds.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dts() dts.Interface {
	return dts.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Eais() eais.Interface {
	return eais.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ecd() ecd.Interface {
	return ecd.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Eci() eci.Interface {
	return eci.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ecp() ecp.Interface {
	return ecp.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ecs() ecs.Interface {
	return ecs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Edas() edas.Interface {
	return edas.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ehpc() ehpc.Interface {
	return ehpc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Eip() eip.Interface {
	return eip.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Eipanycast() eipanycast.Interface {
	return eipanycast.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Elasticsearch() elasticsearch.Interface {
	return elasticsearch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Emr() emr.Interface {
	return emr.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ens() ens.Interface {
	return ens.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ess() ess.Interface {
	return ess.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Event() event.Interface {
	return event.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Express() express.Interface {
	return express.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Fc() fc.Interface {
	return fc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Fnf() fnf.Interface {
	return fnf.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Forward() forward.Interface {
	return forward.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ga() ga.Interface {
	return ga.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Gpdb() gpdb.Interface {
	return gpdb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Graph() graph.Interface {
	return graph.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Havip() havip.Interface {
	return havip.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Hbase() hbase.Interface {
	return hbase.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Hbr() hbr.Interface {
	return hbr.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Image() image.Interface {
	return image.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Imm() imm.Interface {
	return imm.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Imp() imp.Interface {
	return imp.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Instance() instance.Interface {
	return instance.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Iot() iot.Interface {
	return iot.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Key() key.Interface {
	return key.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kms() kms.Interface {
	return kms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kvstore() kvstore.Interface {
	return kvstore.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Launch() launch.Interface {
	return launch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Lindorm() lindorm.Interface {
	return lindorm.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Log() log.Interface {
	return log.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Logtail() logtail.Interface {
	return logtail.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Market() market.Interface {
	return market.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Maxcompute() maxcompute.Interface {
	return maxcompute.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mhub() mhub.Interface {
	return mhub.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mns() mns.Interface {
	return mns.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mongodb() mongodb.Interface {
	return mongodb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Msc() msc.Interface {
	return msc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mse() mse.Interface {
	return mse.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Nas() nas.Interface {
	return nas.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Nat() nat.Interface {
	return nat.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Network() network.Interface {
	return network.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ons() ons.Interface {
	return ons.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Oos() oos.Interface {
	return oos.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Open() open.Interface {
	return open.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Oss() oss.Interface {
	return oss.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ots() ots.Interface {
	return ots.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Polardb() polardb.Interface {
	return polardb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Privatelink() privatelink.Interface {
	return privatelink.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Pvtz() pvtz.Interface {
	return pvtz.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Quick() quick.Interface {
	return quick.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Quotas() quotas.Interface {
	return quotas.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ram() ram.Interface {
	return ram.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Rdc() rdc.Interface {
	return rdc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Rds() rds.Interface {
	return rds.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Reserved() reserved.Interface {
	return reserved.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Resource() resource.Interface {
	return resource.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ros() ros.Interface {
	return ros.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Route() route.Interface {
	return route.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Router() router.Interface {
	return router.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sae() sae.Interface {
	return sae.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sag() sag.Interface {
	return sag.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Scdn() scdn.Interface {
	return scdn.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sddp() sddp.Interface {
	return sddp.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Security() security.Interface {
	return security.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Service() service.Interface {
	return service.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Simple() simple.Interface {
	return simple.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Slb() slb.Interface {
	return slb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Snapshot() snapshot.Interface {
	return snapshot.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Snat() snat.Interface {
	return snat.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ssl() ssl.Interface {
	return ssl.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Subnet() subnet.Interface {
	return subnet.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Tsdb() tsdb.Interface {
	return tsdb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Video() video.Interface {
	return video.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Vod() vod.Interface {
	return vod.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Vpc() vpc.Interface {
	return vpc.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Vpn() vpn.Interface {
	return vpn.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Vswitch() vswitch.Interface {
	return vswitch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Waf() waf.Interface {
	return waf.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Yundun() yundun.Interface {
	return yundun.New(f, f.namespace, f.tweakListOptions)
}
