package webhook

import (
	"encoding/json"
	"time"

	"github.com/chargebee/chargebee-go/v3/models/event"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/businessentitychange"

	"github.com/chargebee/chargebee-go/v3/models/businessentitytransfer"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/taxwithheld"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorder"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorderitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/businessentity"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/differentialprice"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/usagereminderinfo"

	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"

	"github.com/chargebee/chargebee-go/v3/models/rule"

	"github.com/chargebee/chargebee-go/v3/models/paymentschedule"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/paymentvoucher"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"

	"github.com/chargebee/chargebee-go/v3/models/paymentintent"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/contractterm"

	"github.com/chargebee/chargebee-go/v3/models/itemfamily"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/pricevariant"

	"github.com/chargebee/chargebee-go/v3/models/attribute"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/item"

	"github.com/chargebee/chargebee-go/v3/models/ramp"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/token"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/itemfamily"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/purchase"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/itemfamily"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/couponcode"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/pricevariant"

	"github.com/chargebee/chargebee-go/v3/models/attribute"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"

	"github.com/chargebee/chargebee-go/v3/models/ramp"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/recordedpurchase"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/item"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/itemprice"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"

	"github.com/chargebee/chargebee-go/v3/models/contractterm"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/taxwithheld"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/contractterm"

	"github.com/chargebee/chargebee-go/v3/models/paymentschedule"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditem"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/businessentitychange"

	"github.com/chargebee/chargebee-go/v3/models/businessentitytransfer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorder"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelonetimeorderitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/quote"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/paymentschedulescheme"

	"github.com/chargebee/chargebee-go/v3/models/businessentity"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/token"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/rule"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditemprice"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/itemprice"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/virtualbankaccount"

	"github.com/chargebee/chargebee-go/v3/models/paymentschedulescheme"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlementscreateddetail"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/itemprice"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/usagefile"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/differentialprice"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/quote"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/attacheditem"

	"github.com/chargebee/chargebee-go/v3/models/salesorder"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/ramp"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/gift"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/advanceinvoiceschedule"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/token"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/promotionalcredit"

	"github.com/chargebee/chargebee-go/v3/models/ramp"

	"github.com/chargebee/chargebee-go/v3/models/impactedcustomer"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/subscriptionentitlementsupdateddetail"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/attacheditem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/item"

	"github.com/chargebee/chargebee-go/v3/models/coupon"

	"github.com/chargebee/chargebee-go/v3/models/couponset"

	"github.com/chargebee/chargebee-go/v3/models/paymentintent"

	"github.com/chargebee/chargebee-go/v3/models/order"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichanneltransaction"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/taxwithheld"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/pricevariant"

	"github.com/chargebee/chargebee-go/v3/models/attribute"

	"github.com/chargebee/chargebee-go/v3/models/differentialprice"

	"github.com/chargebee/chargebee-go/v3/models/subscription"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/card"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/creditnote"

	"github.com/chargebee/chargebee-go/v3/models/unbilledcharge"

	"github.com/chargebee/chargebee-go/v3/models/rule"

	"github.com/chargebee/chargebee-go/v3/models/contractterm"

	"github.com/chargebee/chargebee-go/v3/models/contractterm"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/feature"

	"github.com/chargebee/chargebee-go/v3/models/metadata"

	"github.com/chargebee/chargebee-go/v3/models/impacteditemprice"

	"github.com/chargebee/chargebee-go/v3/models/impactedsubscription"

	"github.com/chargebee/chargebee-go/v3/models/salesorder"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscription"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitem"

	"github.com/chargebee/chargebee-go/v3/models/omnichannelsubscriptionitemscheduledchange"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/invoice"

	"github.com/chargebee/chargebee-go/v3/models/quote"

	"github.com/chargebee/chargebee-go/v3/models/attacheditem"

	"github.com/chargebee/chargebee-go/v3/models/customer"

	"github.com/chargebee/chargebee-go/v3/models/paymentsource"

	"github.com/chargebee/chargebee-go/v3/models/businessentity"

	"github.com/chargebee/chargebee-go/v3/models/transaction"

	"github.com/chargebee/chargebee-go/v3/models/ramp"
)

// WebhookContentType represents the type of webhook event content
type WebhookContentType string

const (
	WebhookContentTypeSubscriptionPauseScheduled WebhookContentType = "subscription_pause_scheduled"

	WebhookContentTypeCustomerBusinessEntityChanged WebhookContentType = "customer_business_entity_changed"

	WebhookContentTypeSubscriptionAdvanceInvoiceScheduleAdded WebhookContentType = "subscription_advance_invoice_schedule_added"

	WebhookContentTypeGiftExpired WebhookContentType = "gift_expired"

	WebhookContentTypeTaxWithheldDeleted WebhookContentType = "tax_withheld_deleted"

	WebhookContentTypeUnbilledChargesDeleted WebhookContentType = "unbilled_charges_deleted"

	WebhookContentTypeCouponUpdated WebhookContentType = "coupon_updated"

	WebhookContentTypeOmnichannelSubscriptionItemReactivated WebhookContentType = "omnichannel_subscription_item_reactivated"

	WebhookContentTypeOmnichannelSubscriptionItemRenewed WebhookContentType = "omnichannel_subscription_item_renewed"

	WebhookContentTypeUnbilledChargesCreated WebhookContentType = "unbilled_charges_created"

	WebhookContentTypeSubscriptionResumed WebhookContentType = "subscription_resumed"

	WebhookContentTypeOmnichannelOneTimeOrderItemCancelled WebhookContentType = "omnichannel_one_time_order_item_cancelled"

	WebhookContentTypeSubscriptionCancelled WebhookContentType = "subscription_cancelled"

	WebhookContentTypeItemEntitlementsRemoved WebhookContentType = "item_entitlements_removed"

	WebhookContentTypeBusinessEntityCreated WebhookContentType = "business_entity_created"

	WebhookContentTypeCouponSetUpdated WebhookContentType = "coupon_set_updated"

	WebhookContentTypeDifferentialPriceUpdated WebhookContentType = "differential_price_updated"

	WebhookContentTypeOmnichannelSubscriptionItemPaused WebhookContentType = "omnichannel_subscription_item_paused"

	WebhookContentTypeEntitlementOverridesRemoved WebhookContentType = "entitlement_overrides_removed"

	WebhookContentTypeSubscriptionActivatedWithBackdating WebhookContentType = "subscription_activated_with_backdating"

	WebhookContentTypeSubscriptionTrialEndReminder WebhookContentType = "subscription_trial_end_reminder"

	WebhookContentTypeSubscriptionShippingAddressUpdated WebhookContentType = "subscription_shipping_address_updated"

	WebhookContentTypeVoucherCreateFailed WebhookContentType = "voucher_create_failed"

	WebhookContentTypeGiftClaimed WebhookContentType = "gift_claimed"

	WebhookContentTypeCustomerDeleted WebhookContentType = "customer_deleted"

	WebhookContentTypeRefundInitiated WebhookContentType = "refund_initiated"

	WebhookContentTypeInvoiceGeneratedWithBackdating WebhookContentType = "invoice_generated_with_backdating"

	WebhookContentTypeOmnichannelTransactionCreated WebhookContentType = "omnichannel_transaction_created"

	WebhookContentTypeAddUsagesReminder WebhookContentType = "add_usages_reminder"

	WebhookContentTypeVoucherCreated WebhookContentType = "voucher_created"

	WebhookContentTypeRuleUpdated WebhookContentType = "rule_updated"

	WebhookContentTypePaymentSchedulesCreated WebhookContentType = "payment_schedules_created"

	WebhookContentTypeFeatureActivated WebhookContentType = "feature_activated"

	WebhookContentTypePaymentSourceLocallyDeleted WebhookContentType = "payment_source_locally_deleted"

	WebhookContentTypeInvoiceGenerated WebhookContentType = "invoice_generated"

	WebhookContentTypeVoucherExpired WebhookContentType = "voucher_expired"

	WebhookContentTypeAuthorizationSucceeded WebhookContentType = "authorization_succeeded"

	WebhookContentTypeGiftScheduled WebhookContentType = "gift_scheduled"

	WebhookContentTypeSubscriptionChangesScheduled WebhookContentType = "subscription_changes_scheduled"

	WebhookContentTypeSubscriptionChangedWithBackdating WebhookContentType = "subscription_changed_with_backdating"

	WebhookContentTypeOmnichannelSubscriptionItemChanged WebhookContentType = "omnichannel_subscription_item_changed"

	WebhookContentTypeGiftUnclaimed WebhookContentType = "gift_unclaimed"

	WebhookContentTypeVirtualBankAccountAdded WebhookContentType = "virtual_bank_account_added"

	WebhookContentTypePaymentIntentCreated WebhookContentType = "payment_intent_created"

	WebhookContentTypeCreditNoteCreatedWithBackdating WebhookContentType = "credit_note_created_with_backdating"

	WebhookContentTypeContractTermTerminated WebhookContentType = "contract_term_terminated"

	WebhookContentTypeItemFamilyUpdated WebhookContentType = "item_family_updated"

	WebhookContentTypeOrderCreated WebhookContentType = "order_created"

	WebhookContentTypePriceVariantDeleted WebhookContentType = "price_variant_deleted"

	WebhookContentTypeSubscriptionMovementFailed WebhookContentType = "subscription_movement_failed"

	WebhookContentTypeCustomerMovedIn WebhookContentType = "customer_moved_in"

	WebhookContentTypeSubscriptionAdvanceInvoiceScheduleUpdated WebhookContentType = "subscription_advance_invoice_schedule_updated"

	WebhookContentTypeItemDeleted WebhookContentType = "item_deleted"

	WebhookContentTypeSubscriptionRampDrafted WebhookContentType = "subscription_ramp_drafted"

	WebhookContentTypeDunningUpdated WebhookContentType = "dunning_updated"

	WebhookContentTypeItemEntitlementsUpdated WebhookContentType = "item_entitlements_updated"

	WebhookContentTypeTokenConsumed WebhookContentType = "token_consumed"

	WebhookContentTypeHierarchyDeleted WebhookContentType = "hierarchy_deleted"

	WebhookContentTypeSubscriptionCancellationScheduled WebhookContentType = "subscription_cancellation_scheduled"

	WebhookContentTypeSubscriptionRenewed WebhookContentType = "subscription_renewed"

	WebhookContentTypeFeatureUpdated WebhookContentType = "feature_updated"

	WebhookContentTypeFeatureDeleted WebhookContentType = "feature_deleted"

	WebhookContentTypeItemFamilyCreated WebhookContentType = "item_family_created"

	WebhookContentTypeOmnichannelSubscriptionItemScheduledChangeRemoved WebhookContentType = "omnichannel_subscription_item_scheduled_change_removed"

	WebhookContentTypeOmnichannelSubscriptionItemResumed WebhookContentType = "omnichannel_subscription_item_resumed"

	WebhookContentTypePurchaseCreated WebhookContentType = "purchase_created"

	WebhookContentTypeEntitlementOverridesUpdated WebhookContentType = "entitlement_overrides_updated"

	WebhookContentTypeItemFamilyDeleted WebhookContentType = "item_family_deleted"

	WebhookContentTypeSubscriptionResumptionScheduled WebhookContentType = "subscription_resumption_scheduled"

	WebhookContentTypeFeatureReactivated WebhookContentType = "feature_reactivated"

	WebhookContentTypeCouponCodesDeleted WebhookContentType = "coupon_codes_deleted"

	WebhookContentTypeCardExpired WebhookContentType = "card_expired"

	WebhookContentTypeCreditNoteUpdated WebhookContentType = "credit_note_updated"

	WebhookContentTypeOmnichannelSubscriptionItemDowngraded WebhookContentType = "omnichannel_subscription_item_downgraded"

	WebhookContentTypePriceVariantUpdated WebhookContentType = "price_variant_updated"

	WebhookContentTypePromotionalCreditsDeducted WebhookContentType = "promotional_credits_deducted"

	WebhookContentTypeSubscriptionRampApplied WebhookContentType = "subscription_ramp_applied"

	WebhookContentTypeSubscriptionPaused WebhookContentType = "subscription_paused"

	WebhookContentTypeOrderReadyToProcess WebhookContentType = "order_ready_to_process"

	WebhookContentTypeFeatureCreated WebhookContentType = "feature_created"

	WebhookContentTypeTransactionDeleted WebhookContentType = "transaction_deleted"

	WebhookContentTypeCreditNoteCreated WebhookContentType = "credit_note_created"

	WebhookContentTypeOmnichannelSubscriptionItemResubscribed WebhookContentType = "omnichannel_subscription_item_resubscribed"

	WebhookContentTypeRecordPurchaseFailed WebhookContentType = "record_purchase_failed"

	WebhookContentTypeItemCreated WebhookContentType = "item_created"

	WebhookContentTypeTransactionUpdated WebhookContentType = "transaction_updated"

	WebhookContentTypeMrrUpdated WebhookContentType = "mrr_updated"

	WebhookContentTypeUnbilledChargesInvoiced WebhookContentType = "unbilled_charges_invoiced"

	WebhookContentTypeItemPriceUpdated WebhookContentType = "item_price_updated"

	WebhookContentTypeCouponCodesUpdated WebhookContentType = "coupon_codes_updated"

	WebhookContentTypeVirtualBankAccountUpdated WebhookContentType = "virtual_bank_account_updated"

	WebhookContentTypeContractTermCreated WebhookContentType = "contract_term_created"

	WebhookContentTypeSubscriptionChanged WebhookContentType = "subscription_changed"

	WebhookContentTypePaymentFailed WebhookContentType = "payment_failed"

	WebhookContentTypeCreditNoteDeleted WebhookContentType = "credit_note_deleted"

	WebhookContentTypeTaxWithheldRefunded WebhookContentType = "tax_withheld_refunded"

	WebhookContentTypeContractTermCompleted WebhookContentType = "contract_term_completed"

	WebhookContentTypePaymentSchedulesUpdated WebhookContentType = "payment_schedules_updated"

	WebhookContentTypeOmnichannelSubscriptionItemExpired WebhookContentType = "omnichannel_subscription_item_expired"

	WebhookContentTypeCardUpdated WebhookContentType = "card_updated"

	WebhookContentTypeCustomerCreated WebhookContentType = "customer_created"

	WebhookContentTypeSubscriptionRenewalReminder WebhookContentType = "subscription_renewal_reminder"

	WebhookContentTypeOrderDelivered WebhookContentType = "order_delivered"

	WebhookContentTypeOmnichannelSubscriptionItemCancellationScheduled WebhookContentType = "omnichannel_subscription_item_cancellation_scheduled"

	WebhookContentTypeOmnichannelSubscriptionItemGracePeriodExpired WebhookContentType = "omnichannel_subscription_item_grace_period_expired"

	WebhookContentTypeCouponCodesAdded WebhookContentType = "coupon_codes_added"

	WebhookContentTypeGiftCancelled WebhookContentType = "gift_cancelled"

	WebhookContentTypeOrderCancelled WebhookContentType = "order_cancelled"

	WebhookContentTypeCouponDeleted WebhookContentType = "coupon_deleted"

	WebhookContentTypeSubscriptionScheduledChangesRemoved WebhookContentType = "subscription_scheduled_changes_removed"

	WebhookContentTypePendingInvoiceCreated WebhookContentType = "pending_invoice_created"

	WebhookContentTypeEntitlementOverridesAutoRemoved WebhookContentType = "entitlement_overrides_auto_removed"

	WebhookContentTypeOmnichannelSubscriptionItemUpgraded WebhookContentType = "omnichannel_subscription_item_upgraded"

	WebhookContentTypeSubscriptionBusinessEntityChanged WebhookContentType = "subscription_business_entity_changed"

	WebhookContentTypeOmnichannelOneTimeOrderCreated WebhookContentType = "omnichannel_one_time_order_created"

	WebhookContentTypePaymentSourceDeleted WebhookContentType = "payment_source_deleted"

	WebhookContentTypeOmnichannelSubscriptionItemCancelled WebhookContentType = "omnichannel_subscription_item_cancelled"

	WebhookContentTypeQuoteDeleted WebhookContentType = "quote_deleted"

	WebhookContentTypeInvoiceUpdated WebhookContentType = "invoice_updated"

	WebhookContentTypeSubscriptionAdvanceInvoiceScheduleRemoved WebhookContentType = "subscription_advance_invoice_schedule_removed"

	WebhookContentTypeCardDeleted WebhookContentType = "card_deleted"

	WebhookContentTypeOrderReadyToShip WebhookContentType = "order_ready_to_ship"

	WebhookContentTypeSubscriptionMovedOut WebhookContentType = "subscription_moved_out"

	WebhookContentTypePaymentScheduleSchemeCreated WebhookContentType = "payment_schedule_scheme_created"

	WebhookContentTypeBusinessEntityUpdated WebhookContentType = "business_entity_updated"

	WebhookContentTypeSubscriptionScheduledResumptionRemoved WebhookContentType = "subscription_scheduled_resumption_removed"

	WebhookContentTypePaymentInitiated WebhookContentType = "payment_initiated"

	WebhookContentTypeFeatureArchived WebhookContentType = "feature_archived"

	WebhookContentTypeSubscriptionReactivatedWithBackdating WebhookContentType = "subscription_reactivated_with_backdating"

	WebhookContentTypeOmnichannelSubscriptionImported WebhookContentType = "omnichannel_subscription_imported"

	WebhookContentTypeTokenExpired WebhookContentType = "token_expired"

	WebhookContentTypeCardAdded WebhookContentType = "card_added"

	WebhookContentTypeCouponCreated WebhookContentType = "coupon_created"

	WebhookContentTypeRuleDeleted WebhookContentType = "rule_deleted"

	WebhookContentTypeItemPriceEntitlementsUpdated WebhookContentType = "item_price_entitlements_updated"

	WebhookContentTypeItemPriceDeleted WebhookContentType = "item_price_deleted"

	WebhookContentTypeVirtualBankAccountDeleted WebhookContentType = "virtual_bank_account_deleted"

	WebhookContentTypePaymentScheduleSchemeDeleted WebhookContentType = "payment_schedule_scheme_deleted"

	WebhookContentTypeSubscriptionCreated WebhookContentType = "subscription_created"

	WebhookContentTypeSubscriptionEntitlementsCreated WebhookContentType = "subscription_entitlements_created"

	WebhookContentTypeOrderReturned WebhookContentType = "order_returned"

	WebhookContentTypeSubscriptionDeleted WebhookContentType = "subscription_deleted"

	WebhookContentTypePaymentSourceAdded WebhookContentType = "payment_source_added"

	WebhookContentTypeSubscriptionMovedIn WebhookContentType = "subscription_moved_in"

	WebhookContentTypeItemPriceCreated WebhookContentType = "item_price_created"

	WebhookContentTypeSubscriptionScheduledCancellationRemoved WebhookContentType = "subscription_scheduled_cancellation_removed"

	WebhookContentTypePaymentRefunded WebhookContentType = "payment_refunded"

	WebhookContentTypeUsageFileIngested WebhookContentType = "usage_file_ingested"

	WebhookContentTypeOmnichannelSubscriptionMovedIn WebhookContentType = "omnichannel_subscription_moved_in"

	WebhookContentTypeDifferentialPriceCreated WebhookContentType = "differential_price_created"

	WebhookContentTypeTransactionCreated WebhookContentType = "transaction_created"

	WebhookContentTypePaymentSucceeded WebhookContentType = "payment_succeeded"

	WebhookContentTypeSubscriptionCanceledWithBackdating WebhookContentType = "subscription_canceled_with_backdating"

	WebhookContentTypeUnbilledChargesVoided WebhookContentType = "unbilled_charges_voided"

	WebhookContentTypeQuoteCreated WebhookContentType = "quote_created"

	WebhookContentTypeCouponSetDeleted WebhookContentType = "coupon_set_deleted"

	WebhookContentTypeAttachedItemCreated WebhookContentType = "attached_item_created"

	WebhookContentTypeSalesOrderCreated WebhookContentType = "sales_order_created"

	WebhookContentTypeCustomerChanged WebhookContentType = "customer_changed"

	WebhookContentTypeSubscriptionStarted WebhookContentType = "subscription_started"

	WebhookContentTypeSubscriptionActivated WebhookContentType = "subscription_activated"

	WebhookContentTypePaymentSourceExpiring WebhookContentType = "payment_source_expiring"

	WebhookContentTypeSubscriptionReactivated WebhookContentType = "subscription_reactivated"

	WebhookContentTypeOrderUpdated WebhookContentType = "order_updated"

	WebhookContentTypeSubscriptionScheduledPauseRemoved WebhookContentType = "subscription_scheduled_pause_removed"

	WebhookContentTypeSubscriptionCancellationReminder WebhookContentType = "subscription_cancellation_reminder"

	WebhookContentTypeSubscriptionCreatedWithBackdating WebhookContentType = "subscription_created_with_backdating"

	WebhookContentTypeSubscriptionRampCreated WebhookContentType = "subscription_ramp_created"

	WebhookContentTypeOrderDeleted WebhookContentType = "order_deleted"

	WebhookContentTypeOmnichannelSubscriptionItemPauseScheduled WebhookContentType = "omnichannel_subscription_item_pause_scheduled"

	WebhookContentTypeGiftUpdated WebhookContentType = "gift_updated"

	WebhookContentTypeSubscriptionTrialExtended WebhookContentType = "subscription_trial_extended"

	WebhookContentTypeOmnichannelSubscriptionItemGracePeriodStarted WebhookContentType = "omnichannel_subscription_item_grace_period_started"

	WebhookContentTypeCardExpiryReminder WebhookContentType = "card_expiry_reminder"

	WebhookContentTypeTokenCreated WebhookContentType = "token_created"

	WebhookContentTypePromotionalCreditsAdded WebhookContentType = "promotional_credits_added"

	WebhookContentTypeSubscriptionRampUpdated WebhookContentType = "subscription_ramp_updated"

	WebhookContentTypeCustomerEntitlementsUpdated WebhookContentType = "customer_entitlements_updated"

	WebhookContentTypePaymentSourceExpired WebhookContentType = "payment_source_expired"

	WebhookContentTypeCustomerMovedOut WebhookContentType = "customer_moved_out"

	WebhookContentTypeSubscriptionEntitlementsUpdated WebhookContentType = "subscription_entitlements_updated"

	WebhookContentTypeOmnichannelSubscriptionItemDunningExpired WebhookContentType = "omnichannel_subscription_item_dunning_expired"

	WebhookContentTypeHierarchyCreated WebhookContentType = "hierarchy_created"

	WebhookContentTypeAttachedItemDeleted WebhookContentType = "attached_item_deleted"

	WebhookContentTypeOmnichannelSubscriptionItemScheduledCancellationRemoved WebhookContentType = "omnichannel_subscription_item_scheduled_cancellation_removed"

	WebhookContentTypeItemUpdated WebhookContentType = "item_updated"

	WebhookContentTypeCouponSetCreated WebhookContentType = "coupon_set_created"

	WebhookContentTypePaymentIntentUpdated WebhookContentType = "payment_intent_updated"

	WebhookContentTypeOrderResent WebhookContentType = "order_resent"

	WebhookContentTypeOmnichannelSubscriptionCreated WebhookContentType = "omnichannel_subscription_created"

	WebhookContentTypeTaxWithheldRecorded WebhookContentType = "tax_withheld_recorded"

	WebhookContentTypePriceVariantCreated WebhookContentType = "price_variant_created"

	WebhookContentTypeDifferentialPriceDeleted WebhookContentType = "differential_price_deleted"

	WebhookContentTypeSubscriptionItemsRenewed WebhookContentType = "subscription_items_renewed"

	WebhookContentTypeRuleCreated WebhookContentType = "rule_created"

	WebhookContentTypeContractTermCancelled WebhookContentType = "contract_term_cancelled"

	WebhookContentTypeContractTermRenewed WebhookContentType = "contract_term_renewed"

	WebhookContentTypeInvoiceDeleted WebhookContentType = "invoice_deleted"

	WebhookContentTypeItemPriceEntitlementsRemoved WebhookContentType = "item_price_entitlements_removed"

	WebhookContentTypeSalesOrderUpdated WebhookContentType = "sales_order_updated"

	WebhookContentTypeOmnichannelSubscriptionItemDunningStarted WebhookContentType = "omnichannel_subscription_item_dunning_started"

	WebhookContentTypeOmnichannelSubscriptionItemChangeScheduled WebhookContentType = "omnichannel_subscription_item_change_scheduled"

	WebhookContentTypePendingInvoiceUpdated WebhookContentType = "pending_invoice_updated"

	WebhookContentTypeQuoteUpdated WebhookContentType = "quote_updated"

	WebhookContentTypeAttachedItemUpdated WebhookContentType = "attached_item_updated"

	WebhookContentTypePaymentSourceUpdated WebhookContentType = "payment_source_updated"

	WebhookContentTypeBusinessEntityDeleted WebhookContentType = "business_entity_deleted"

	WebhookContentTypeAuthorizationVoided WebhookContentType = "authorization_voided"

	WebhookContentTypeSubscriptionRampDeleted WebhookContentType = "subscription_ramp_deleted"
)

// WebhookEvent represents a Chargebee webhook event with typed content
type WebhookEvent[T WebhookContentType] struct {
	Id                   string          `json:"id"`
	OccurredAt           int64           `json:"occurred_at"`
	Source               string          `json:"source"`
	User                 string          `json:"user"`
	WebhookStatus        string          `json:"webhook_status"`
	WebhookFailureReason string          `json:"webhook_failure_reason"`
	EventType            string          `json:"event_type"`
	ApiVersion           string          `json:"api_version"`
	Content              json.RawMessage `json:"content"`
	OriginUser           string          `json:"origin_user"`
	Object               string          `json:"object"`
}

// GetOccurredAtTime returns the occurred_at timestamp as a time.Time
func (e *WebhookEvent[T]) GetOccurredAtTime() time.Time {
	return time.Unix(e.OccurredAt, 0)
}

// Event content structures for each webhook type

type SubscriptionPauseScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type CustomerBusinessEntityChangedContent struct {
	BusinessEntityChange *businessentitychange.BusinessEntityChange `json:"business_entity_change,omitempty"`

	BusinessEntityTransfer *businessentitytransfer.BusinessEntityTransfer `json:"business_entity_transfer,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleAddedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type GiftExpiredContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type TaxWithheldDeletedContent struct {
	TaxWithheld *taxwithheld.TaxWithheld `json:"tax_withheld,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type UnbilledChargesDeletedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type CouponUpdatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`
}

type OmnichannelSubscriptionItemReactivatedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemRenewedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type UnbilledChargesCreatedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionResumedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelOneTimeOrderItemCancelledContent struct {
	OmnichannelOneTimeOrder *omnichannelonetimeorder.OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`

	OmnichannelOneTimeOrderItem *omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_item,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionCancelledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type ItemEntitlementsRemovedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type BusinessEntityCreatedContent struct {
	BusinessEntity *businessentity.BusinessEntity `json:"business_entity,omitempty"`
}

type CouponSetUpdatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type DifferentialPriceUpdatedContent struct {
	DifferentialPrice *differentialprice.DifferentialPrice `json:"differential_price,omitempty"`
}

type OmnichannelSubscriptionItemPausedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type EntitlementOverridesRemovedContent struct {
	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type SubscriptionActivatedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionTrialEndReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionShippingAddressUpdatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type VoucherCreateFailedContent struct {
	PaymentVoucher *paymentvoucher.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type GiftClaimedContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type CustomerDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type RefundInitiatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type InvoiceGeneratedWithBackdatingContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type OmnichannelTransactionCreatedContent struct {
	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`
}

type AddUsagesReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	UsageReminderInfo *usagereminderinfo.UsageReminderInfo `json:"usage_reminder_info,omitempty"`
}

type VoucherCreatedContent struct {
	PaymentVoucher *paymentvoucher.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type RuleUpdatedContent struct {
	Rule *rule.Rule `json:"rule,omitempty"`
}

type PaymentSchedulesCreatedContent struct {
	PaymentSchedule *paymentschedule.PaymentSchedule `json:"payment_schedule,omitempty"`
}

type FeatureActivatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type PaymentSourceLocallyDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type InvoiceGeneratedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type VoucherExpiredContent struct {
	PaymentVoucher *paymentvoucher.PaymentVoucher `json:"payment_voucher,omitempty"`
}

type AuthorizationSucceededContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type GiftScheduledContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type SubscriptionChangesScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionChangedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelSubscriptionItemChangedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type GiftUnclaimedContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type VirtualBankAccountAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	VirtualBankAccount *virtualbankaccount.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type PaymentIntentCreatedContent struct {
	PaymentIntent *paymentintent.PaymentIntent `json:"payment_intent,omitempty"`
}

type CreditNoteCreatedWithBackdatingContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type ContractTermTerminatedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type ItemFamilyUpdatedContent struct {
	ItemFamily *itemfamily.ItemFamily `json:"item_family,omitempty"`
}

type OrderCreatedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type PriceVariantDeletedContent struct {
	PriceVariant *pricevariant.PriceVariant `json:"price_variant,omitempty"`

	Attribute *attribute.Attribute `json:"attribute,omitempty"`
}

type SubscriptionMovementFailedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type CustomerMovedInContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleUpdatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type ItemDeletedContent struct {
	Item *item.Item `json:"item,omitempty"`
}

type SubscriptionRampDraftedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type DunningUpdatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type ItemEntitlementsUpdatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type TokenConsumedContent struct {
	Token *token.Token `json:"token,omitempty"`
}

type HierarchyDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionCancellationScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionRenewedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type FeatureUpdatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type FeatureDeletedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type ItemFamilyCreatedContent struct {
	ItemFamily *itemfamily.ItemFamily `json:"item_family,omitempty"`
}

type OmnichannelSubscriptionItemScheduledChangeRemovedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemResumedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PurchaseCreatedContent struct {
	Purchase *purchase.Purchase `json:"purchase,omitempty"`
}

type EntitlementOverridesUpdatedContent struct {
	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type ItemFamilyDeletedContent struct {
	ItemFamily *itemfamily.ItemFamily `json:"item_family,omitempty"`
}

type SubscriptionResumptionScheduledContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type FeatureReactivatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type CouponCodesDeletedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`

	CouponCode *couponcode.CouponCode `json:"coupon_code,omitempty"`
}

type CardExpiredContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CreditNoteUpdatedContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type OmnichannelSubscriptionItemDowngradedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PriceVariantUpdatedContent struct {
	PriceVariant *pricevariant.PriceVariant `json:"price_variant,omitempty"`

	Attribute *attribute.Attribute `json:"attribute,omitempty"`
}

type PromotionalCreditsDeductedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PromotionalCredit *promotionalcredit.PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SubscriptionRampAppliedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type SubscriptionPausedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OrderReadyToProcessContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type FeatureCreatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type TransactionDeletedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type CreditNoteCreatedContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type OmnichannelSubscriptionItemResubscribedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type RecordPurchaseFailedContent struct {
	RecordedPurchase *recordedpurchase.RecordedPurchase `json:"recorded_purchase,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type ItemCreatedContent struct {
	Item *item.Item `json:"item,omitempty"`
}

type TransactionUpdatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type MrrUpdatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type UnbilledChargesInvoicedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type ItemPriceUpdatedContent struct {
	ItemPrice *itemprice.ItemPrice `json:"item_price,omitempty"`
}

type CouponCodesUpdatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type VirtualBankAccountUpdatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	VirtualBankAccount *virtualbankaccount.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type ContractTermCreatedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type SubscriptionChangedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type PaymentFailedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CreditNoteDeletedContent struct {
	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type TaxWithheldRefundedContent struct {
	TaxWithheld *taxwithheld.TaxWithheld `json:"tax_withheld,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type ContractTermCompletedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type PaymentSchedulesUpdatedContent struct {
	PaymentSchedule *paymentschedule.PaymentSchedule `json:"payment_schedule,omitempty"`
}

type OmnichannelSubscriptionItemExpiredContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type CardUpdatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CustomerCreatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionRenewalReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type OrderDeliveredContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemCancellationScheduledContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemGracePeriodExpiredContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type CouponCodesAddedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type GiftCancelledContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type OrderCancelledContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type CouponDeletedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`
}

type SubscriptionScheduledChangesRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PendingInvoiceCreatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type EntitlementOverridesAutoRemovedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItem *impacteditem.ImpactedItem `json:"impacted_item,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type OmnichannelSubscriptionItemUpgradedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type SubscriptionBusinessEntityChangedContent struct {
	BusinessEntityChange *businessentitychange.BusinessEntityChange `json:"business_entity_change,omitempty"`

	BusinessEntityTransfer *businessentitytransfer.BusinessEntityTransfer `json:"business_entity_transfer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type OmnichannelOneTimeOrderCreatedContent struct {
	OmnichannelOneTimeOrder *omnichannelonetimeorder.OmnichannelOneTimeOrder `json:"omnichannel_one_time_order,omitempty"`

	OmnichannelOneTimeOrderItem *omnichannelonetimeorderitem.OmnichannelOneTimeOrderItem `json:"omnichannel_one_time_order_item,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PaymentSourceDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type OmnichannelSubscriptionItemCancelledContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type QuoteDeletedContent struct {
	Quote *quote.Quote `json:"quote,omitempty"`
}

type InvoiceUpdatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type SubscriptionAdvanceInvoiceScheduleRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type CardDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type OrderReadyToShipContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type SubscriptionMovedOutContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type PaymentScheduleSchemeCreatedContent struct {
	PaymentScheduleScheme *paymentschedulescheme.PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type BusinessEntityUpdatedContent struct {
	BusinessEntity *businessentity.BusinessEntity `json:"business_entity,omitempty"`
}

type SubscriptionScheduledResumptionRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentInitiatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type FeatureArchivedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

type SubscriptionReactivatedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OmnichannelSubscriptionImportedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type TokenExpiredContent struct {
	Token *token.Token `json:"token,omitempty"`
}

type CardAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type CouponCreatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`
}

type RuleDeletedContent struct {
	Rule *rule.Rule `json:"rule,omitempty"`
}

type ItemPriceEntitlementsUpdatedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItemPrice *impacteditemprice.ImpactedItemPrice `json:"impacted_item_price,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type ItemPriceDeletedContent struct {
	ItemPrice *itemprice.ItemPrice `json:"item_price,omitempty"`
}

type VirtualBankAccountDeletedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	VirtualBankAccount *virtualbankaccount.VirtualBankAccount `json:"virtual_bank_account,omitempty"`
}

type PaymentScheduleSchemeDeletedContent struct {
	PaymentScheduleScheme *paymentschedulescheme.PaymentScheduleScheme `json:"payment_schedule_scheme,omitempty"`
}

type SubscriptionCreatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionEntitlementsCreatedContent struct {
	SubscriptionEntitlementsCreatedDetail *subscriptionentitlementscreateddetail.SubscriptionEntitlementsCreatedDetail `json:"subscription_entitlements_created_detail,omitempty"`
}

type OrderReturnedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type SubscriptionDeletedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentSourceAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type SubscriptionMovedInContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`
}

type ItemPriceCreatedContent struct {
	ItemPrice *itemprice.ItemPrice `json:"item_price,omitempty"`
}

type SubscriptionScheduledCancellationRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type PaymentRefundedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type UsageFileIngestedContent struct {
	UsageFile *usagefile.UsageFile `json:"usage_file,omitempty"`
}

type OmnichannelSubscriptionMovedInContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type DifferentialPriceCreatedContent struct {
	DifferentialPrice *differentialprice.DifferentialPrice `json:"differential_price,omitempty"`
}

type TransactionCreatedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type PaymentSucceededContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionCanceledWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type UnbilledChargesVoidedContent struct {
	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type QuoteCreatedContent struct {
	Quote *quote.Quote `json:"quote,omitempty"`
}

type CouponSetDeletedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type AttachedItemCreatedContent struct {
	AttachedItem *attacheditem.AttachedItem `json:"attached_item,omitempty"`
}

type SalesOrderCreatedContent struct {
	SalesOrder *salesorder.SalesOrder `json:"sales_order,omitempty"`
}

type CustomerChangedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionStartedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type SubscriptionActivatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type PaymentSourceExpiringContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type SubscriptionReactivatedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type OrderUpdatedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type SubscriptionScheduledPauseRemovedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionCancellationReminderContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type SubscriptionCreatedWithBackdatingContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type SubscriptionRampCreatedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type OrderDeletedContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionItemPauseScheduledContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type GiftUpdatedContent struct {
	Gift *gift.Gift `json:"gift,omitempty"`
}

type SubscriptionTrialExtendedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	AdvanceInvoiceSchedule *advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedule,omitempty"`
}

type OmnichannelSubscriptionItemGracePeriodStartedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type CardExpiryReminderContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type TokenCreatedContent struct {
	Token *token.Token `json:"token,omitempty"`
}

type PromotionalCreditsAddedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PromotionalCredit *promotionalcredit.PromotionalCredit `json:"promotional_credit,omitempty"`
}

type SubscriptionRampUpdatedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

type CustomerEntitlementsUpdatedContent struct {
	ImpactedCustomer *impactedcustomer.ImpactedCustomer `json:"impacted_customer,omitempty"`
}

type PaymentSourceExpiredContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type CustomerMovedOutContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`
}

type SubscriptionEntitlementsUpdatedContent struct {
	SubscriptionEntitlementsUpdatedDetail *subscriptionentitlementsupdateddetail.SubscriptionEntitlementsUpdatedDetail `json:"subscription_entitlements_updated_detail,omitempty"`
}

type OmnichannelSubscriptionItemDunningExpiredContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type HierarchyCreatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`
}

type AttachedItemDeletedContent struct {
	AttachedItem *attacheditem.AttachedItem `json:"attached_item,omitempty"`
}

type OmnichannelSubscriptionItemScheduledCancellationRemovedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type ItemUpdatedContent struct {
	Item *item.Item `json:"item,omitempty"`
}

type CouponSetCreatedContent struct {
	Coupon *coupon.Coupon `json:"coupon,omitempty"`

	CouponSet *couponset.CouponSet `json:"coupon_set,omitempty"`
}

type PaymentIntentUpdatedContent struct {
	PaymentIntent *paymentintent.PaymentIntent `json:"payment_intent,omitempty"`
}

type OrderResentContent struct {
	Order *order.Order `json:"order,omitempty"`
}

type OmnichannelSubscriptionCreatedContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelTransaction *omnichanneltransaction.OmnichannelTransaction `json:"omnichannel_transaction,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type TaxWithheldRecordedContent struct {
	TaxWithheld *taxwithheld.TaxWithheld `json:"tax_withheld,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`
}

type PriceVariantCreatedContent struct {
	PriceVariant *pricevariant.PriceVariant `json:"price_variant,omitempty"`

	Attribute *attribute.Attribute `json:"attribute,omitempty"`
}

type DifferentialPriceDeletedContent struct {
	DifferentialPrice *differentialprice.DifferentialPrice `json:"differential_price,omitempty"`
}

type SubscriptionItemsRenewedContent struct {
	Subscription *subscription.Subscription `json:"subscription,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`

	Card *card.Card `json:"card,omitempty"`

	Invoice *invoice.Invoice `json:"invoice,omitempty"`

	CreditNote *creditnote.CreditNote `json:"credit_note,omitempty"`

	UnbilledCharge *unbilledcharge.UnbilledCharge `json:"unbilled_charge,omitempty"`
}

type RuleCreatedContent struct {
	Rule *rule.Rule `json:"rule,omitempty"`
}

type ContractTermCancelledContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type ContractTermRenewedContent struct {
	ContractTerm *contractterm.ContractTerm `json:"contract_term,omitempty"`
}

type InvoiceDeletedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type ItemPriceEntitlementsRemovedContent struct {
	Feature *feature.Feature `json:"feature,omitempty"`

	Metadata *metadata.Metadata `json:"metadata,omitempty"`

	ImpactedItemPrice *impacteditemprice.ImpactedItemPrice `json:"impacted_item_price,omitempty"`

	ImpactedSubscription *impactedsubscription.ImpactedSubscription `json:"impacted_subscription,omitempty"`
}

type SalesOrderUpdatedContent struct {
	SalesOrder *salesorder.SalesOrder `json:"sales_order,omitempty"`
}

type OmnichannelSubscriptionItemDunningStartedContent struct {
	OmnichannelSubscription *omnichannelsubscription.OmnichannelSubscription `json:"omnichannel_subscription,omitempty"`

	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type OmnichannelSubscriptionItemChangeScheduledContent struct {
	OmnichannelSubscriptionItem *omnichannelsubscriptionitem.OmnichannelSubscriptionItem `json:"omnichannel_subscription_item,omitempty"`

	OmnichannelSubscriptionItemScheduledChange *omnichannelsubscriptionitemscheduledchange.OmnichannelSubscriptionItemScheduledChange `json:"omnichannel_subscription_item_scheduled_change,omitempty"`

	Customer *customer.Customer `json:"customer,omitempty"`
}

type PendingInvoiceUpdatedContent struct {
	Invoice *invoice.Invoice `json:"invoice,omitempty"`
}

type QuoteUpdatedContent struct {
	Quote *quote.Quote `json:"quote,omitempty"`
}

type AttachedItemUpdatedContent struct {
	AttachedItem *attacheditem.AttachedItem `json:"attached_item,omitempty"`
}

type PaymentSourceUpdatedContent struct {
	Customer *customer.Customer `json:"customer,omitempty"`

	PaymentSource *paymentsource.PaymentSource `json:"payment_source,omitempty"`
}

type BusinessEntityDeletedContent struct {
	BusinessEntity *businessentity.BusinessEntity `json:"business_entity,omitempty"`
}

type AuthorizationVoidedContent struct {
	Transaction *transaction.Transaction `json:"transaction,omitempty"`
}

type SubscriptionRampDeletedContent struct {
	Ramp *ramp.Ramp `json:"ramp,omitempty"`
}

// WebhookEventInterface defines the common interface for all webhook events
type WebhookEventInterface interface {
	GetEventType() string
	GetEventId() string
	GetOccurredAt() time.Time
}

// Generated event types for each webhook event

// SubscriptionPauseScheduledEvent represents a subscription_pause_scheduled webhook event
type SubscriptionPauseScheduledEvent struct {
	Id         string                             `json:"id"`
	OccurredAt time.Time                          `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *SubscriptionPauseScheduledContent `json:"content"`
}

func (e *SubscriptionPauseScheduledEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionPauseScheduledEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionPauseScheduledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerBusinessEntityChangedEvent represents a customer_business_entity_changed webhook event
type CustomerBusinessEntityChangedEvent struct {
	Id         string                                `json:"id"`
	OccurredAt time.Time                             `json:"occurred_at"`
	EventType  string                                `json:"event_type"`
	Content    *CustomerBusinessEntityChangedContent `json:"content"`
}

func (e *CustomerBusinessEntityChangedEvent) GetEventType() string     { return e.EventType }
func (e *CustomerBusinessEntityChangedEvent) GetEventId() string       { return e.Id }
func (e *CustomerBusinessEntityChangedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionAdvanceInvoiceScheduleAddedEvent represents a subscription_advance_invoice_schedule_added webhook event
type SubscriptionAdvanceInvoiceScheduleAddedEvent struct {
	Id         string                                          `json:"id"`
	OccurredAt time.Time                                       `json:"occurred_at"`
	EventType  string                                          `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleAddedContent `json:"content"`
}

func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionAdvanceInvoiceScheduleAddedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// GiftExpiredEvent represents a gift_expired webhook event
type GiftExpiredEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftExpiredContent `json:"content"`
}

func (e *GiftExpiredEvent) GetEventType() string     { return e.EventType }
func (e *GiftExpiredEvent) GetEventId() string       { return e.Id }
func (e *GiftExpiredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TaxWithheldDeletedEvent represents a tax_withheld_deleted webhook event
type TaxWithheldDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TaxWithheldDeletedContent `json:"content"`
}

func (e *TaxWithheldDeletedEvent) GetEventType() string     { return e.EventType }
func (e *TaxWithheldDeletedEvent) GetEventId() string       { return e.Id }
func (e *TaxWithheldDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// UnbilledChargesDeletedEvent represents a unbilled_charges_deleted webhook event
type UnbilledChargesDeletedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt time.Time                      `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *UnbilledChargesDeletedContent `json:"content"`
}

func (e *UnbilledChargesDeletedEvent) GetEventType() string     { return e.EventType }
func (e *UnbilledChargesDeletedEvent) GetEventId() string       { return e.Id }
func (e *UnbilledChargesDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponUpdatedEvent represents a coupon_updated webhook event
type CouponUpdatedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponUpdatedContent `json:"content"`
}

func (e *CouponUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *CouponUpdatedEvent) GetEventId() string       { return e.Id }
func (e *CouponUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemReactivatedEvent represents a omnichannel_subscription_item_reactivated webhook event
type OmnichannelSubscriptionItemReactivatedEvent struct {
	Id         string                                         `json:"id"`
	OccurredAt time.Time                                      `json:"occurred_at"`
	EventType  string                                         `json:"event_type"`
	Content    *OmnichannelSubscriptionItemReactivatedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemReactivatedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemReactivatedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemReactivatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemRenewedEvent represents a omnichannel_subscription_item_renewed webhook event
type OmnichannelSubscriptionItemRenewedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt time.Time                                  `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemRenewedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemRenewedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemRenewedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemRenewedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// UnbilledChargesCreatedEvent represents a unbilled_charges_created webhook event
type UnbilledChargesCreatedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt time.Time                      `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *UnbilledChargesCreatedContent `json:"content"`
}

func (e *UnbilledChargesCreatedEvent) GetEventType() string     { return e.EventType }
func (e *UnbilledChargesCreatedEvent) GetEventId() string       { return e.Id }
func (e *UnbilledChargesCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionResumedEvent represents a subscription_resumed webhook event
type SubscriptionResumedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionResumedContent `json:"content"`
}

func (e *SubscriptionResumedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionResumedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionResumedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelOneTimeOrderItemCancelledEvent represents a omnichannel_one_time_order_item_cancelled webhook event
type OmnichannelOneTimeOrderItemCancelledEvent struct {
	Id         string                                       `json:"id"`
	OccurredAt time.Time                                    `json:"occurred_at"`
	EventType  string                                       `json:"event_type"`
	Content    *OmnichannelOneTimeOrderItemCancelledContent `json:"content"`
}

func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelOneTimeOrderItemCancelledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionCancelledEvent represents a subscription_cancelled webhook event
type SubscriptionCancelledEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *SubscriptionCancelledContent `json:"content"`
}

func (e *SubscriptionCancelledEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionCancelledEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionCancelledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemEntitlementsRemovedEvent represents a item_entitlements_removed webhook event
type ItemEntitlementsRemovedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *ItemEntitlementsRemovedContent `json:"content"`
}

func (e *ItemEntitlementsRemovedEvent) GetEventType() string     { return e.EventType }
func (e *ItemEntitlementsRemovedEvent) GetEventId() string       { return e.Id }
func (e *ItemEntitlementsRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// BusinessEntityCreatedEvent represents a business_entity_created webhook event
type BusinessEntityCreatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityCreatedContent `json:"content"`
}

func (e *BusinessEntityCreatedEvent) GetEventType() string     { return e.EventType }
func (e *BusinessEntityCreatedEvent) GetEventId() string       { return e.Id }
func (e *BusinessEntityCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponSetUpdatedEvent represents a coupon_set_updated webhook event
type CouponSetUpdatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetUpdatedContent `json:"content"`
}

func (e *CouponSetUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *CouponSetUpdatedEvent) GetEventId() string       { return e.Id }
func (e *CouponSetUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// DifferentialPriceUpdatedEvent represents a differential_price_updated webhook event
type DifferentialPriceUpdatedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt time.Time                        `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceUpdatedContent `json:"content"`
}

func (e *DifferentialPriceUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *DifferentialPriceUpdatedEvent) GetEventId() string       { return e.Id }
func (e *DifferentialPriceUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemPausedEvent represents a omnichannel_subscription_item_paused webhook event
type OmnichannelSubscriptionItemPausedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt time.Time                                 `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *OmnichannelSubscriptionItemPausedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemPausedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemPausedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemPausedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// EntitlementOverridesRemovedEvent represents a entitlement_overrides_removed webhook event
type EntitlementOverridesRemovedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt time.Time                           `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *EntitlementOverridesRemovedContent `json:"content"`
}

func (e *EntitlementOverridesRemovedEvent) GetEventType() string     { return e.EventType }
func (e *EntitlementOverridesRemovedEvent) GetEventId() string       { return e.Id }
func (e *EntitlementOverridesRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionActivatedWithBackdatingEvent represents a subscription_activated_with_backdating webhook event
type SubscriptionActivatedWithBackdatingEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt time.Time                                   `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *SubscriptionActivatedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionActivatedWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionActivatedWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionActivatedWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionTrialEndReminderEvent represents a subscription_trial_end_reminder webhook event
type SubscriptionTrialEndReminderEvent struct {
	Id         string                               `json:"id"`
	OccurredAt time.Time                            `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *SubscriptionTrialEndReminderContent `json:"content"`
}

func (e *SubscriptionTrialEndReminderEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionTrialEndReminderEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionTrialEndReminderEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionShippingAddressUpdatedEvent represents a subscription_shipping_address_updated webhook event
type SubscriptionShippingAddressUpdatedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt time.Time                                  `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *SubscriptionShippingAddressUpdatedContent `json:"content"`
}

func (e *SubscriptionShippingAddressUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionShippingAddressUpdatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionShippingAddressUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// VoucherCreateFailedEvent represents a voucher_create_failed webhook event
type VoucherCreateFailedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *VoucherCreateFailedContent `json:"content"`
}

func (e *VoucherCreateFailedEvent) GetEventType() string     { return e.EventType }
func (e *VoucherCreateFailedEvent) GetEventId() string       { return e.Id }
func (e *VoucherCreateFailedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// GiftClaimedEvent represents a gift_claimed webhook event
type GiftClaimedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftClaimedContent `json:"content"`
}

func (e *GiftClaimedEvent) GetEventType() string     { return e.EventType }
func (e *GiftClaimedEvent) GetEventId() string       { return e.Id }
func (e *GiftClaimedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerDeletedEvent represents a customer_deleted webhook event
type CustomerDeletedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerDeletedContent `json:"content"`
}

func (e *CustomerDeletedEvent) GetEventType() string     { return e.EventType }
func (e *CustomerDeletedEvent) GetEventId() string       { return e.Id }
func (e *CustomerDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// RefundInitiatedEvent represents a refund_initiated webhook event
type RefundInitiatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *RefundInitiatedContent `json:"content"`
}

func (e *RefundInitiatedEvent) GetEventType() string     { return e.EventType }
func (e *RefundInitiatedEvent) GetEventId() string       { return e.Id }
func (e *RefundInitiatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// InvoiceGeneratedWithBackdatingEvent represents a invoice_generated_with_backdating webhook event
type InvoiceGeneratedWithBackdatingEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt time.Time                              `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *InvoiceGeneratedWithBackdatingContent `json:"content"`
}

func (e *InvoiceGeneratedWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *InvoiceGeneratedWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *InvoiceGeneratedWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelTransactionCreatedEvent represents a omnichannel_transaction_created webhook event
type OmnichannelTransactionCreatedEvent struct {
	Id         string                                `json:"id"`
	OccurredAt time.Time                             `json:"occurred_at"`
	EventType  string                                `json:"event_type"`
	Content    *OmnichannelTransactionCreatedContent `json:"content"`
}

func (e *OmnichannelTransactionCreatedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelTransactionCreatedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelTransactionCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// AddUsagesReminderEvent represents a add_usages_reminder webhook event
type AddUsagesReminderEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *AddUsagesReminderContent `json:"content"`
}

func (e *AddUsagesReminderEvent) GetEventType() string     { return e.EventType }
func (e *AddUsagesReminderEvent) GetEventId() string       { return e.Id }
func (e *AddUsagesReminderEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// VoucherCreatedEvent represents a voucher_created webhook event
type VoucherCreatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *VoucherCreatedContent `json:"content"`
}

func (e *VoucherCreatedEvent) GetEventType() string     { return e.EventType }
func (e *VoucherCreatedEvent) GetEventId() string       { return e.Id }
func (e *VoucherCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// RuleUpdatedEvent represents a rule_updated webhook event
type RuleUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleUpdatedContent `json:"content"`
}

func (e *RuleUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *RuleUpdatedEvent) GetEventId() string       { return e.Id }
func (e *RuleUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSchedulesCreatedEvent represents a payment_schedules_created webhook event
type PaymentSchedulesCreatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PaymentSchedulesCreatedContent `json:"content"`
}

func (e *PaymentSchedulesCreatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSchedulesCreatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentSchedulesCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// FeatureActivatedEvent represents a feature_activated webhook event
type FeatureActivatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *FeatureActivatedContent `json:"content"`
}

func (e *FeatureActivatedEvent) GetEventType() string     { return e.EventType }
func (e *FeatureActivatedEvent) GetEventId() string       { return e.Id }
func (e *FeatureActivatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSourceLocallyDeletedEvent represents a payment_source_locally_deleted webhook event
type PaymentSourceLocallyDeletedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt time.Time                           `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *PaymentSourceLocallyDeletedContent `json:"content"`
}

func (e *PaymentSourceLocallyDeletedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSourceLocallyDeletedEvent) GetEventId() string       { return e.Id }
func (e *PaymentSourceLocallyDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// InvoiceGeneratedEvent represents a invoice_generated webhook event
type InvoiceGeneratedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *InvoiceGeneratedContent `json:"content"`
}

func (e *InvoiceGeneratedEvent) GetEventType() string     { return e.EventType }
func (e *InvoiceGeneratedEvent) GetEventId() string       { return e.Id }
func (e *InvoiceGeneratedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// VoucherExpiredEvent represents a voucher_expired webhook event
type VoucherExpiredEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *VoucherExpiredContent `json:"content"`
}

func (e *VoucherExpiredEvent) GetEventType() string     { return e.EventType }
func (e *VoucherExpiredEvent) GetEventId() string       { return e.Id }
func (e *VoucherExpiredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// AuthorizationSucceededEvent represents a authorization_succeeded webhook event
type AuthorizationSucceededEvent struct {
	Id         string                         `json:"id"`
	OccurredAt time.Time                      `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *AuthorizationSucceededContent `json:"content"`
}

func (e *AuthorizationSucceededEvent) GetEventType() string     { return e.EventType }
func (e *AuthorizationSucceededEvent) GetEventId() string       { return e.Id }
func (e *AuthorizationSucceededEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// GiftScheduledEvent represents a gift_scheduled webhook event
type GiftScheduledEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftScheduledContent `json:"content"`
}

func (e *GiftScheduledEvent) GetEventType() string     { return e.EventType }
func (e *GiftScheduledEvent) GetEventId() string       { return e.Id }
func (e *GiftScheduledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionChangesScheduledEvent represents a subscription_changes_scheduled webhook event
type SubscriptionChangesScheduledEvent struct {
	Id         string                               `json:"id"`
	OccurredAt time.Time                            `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *SubscriptionChangesScheduledContent `json:"content"`
}

func (e *SubscriptionChangesScheduledEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionChangesScheduledEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionChangesScheduledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionChangedWithBackdatingEvent represents a subscription_changed_with_backdating webhook event
type SubscriptionChangedWithBackdatingEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt time.Time                                 `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionChangedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionChangedWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionChangedWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionChangedWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemChangedEvent represents a omnichannel_subscription_item_changed webhook event
type OmnichannelSubscriptionItemChangedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt time.Time                                  `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemChangedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemChangedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemChangedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemChangedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// GiftUnclaimedEvent represents a gift_unclaimed webhook event
type GiftUnclaimedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftUnclaimedContent `json:"content"`
}

func (e *GiftUnclaimedEvent) GetEventType() string     { return e.EventType }
func (e *GiftUnclaimedEvent) GetEventId() string       { return e.Id }
func (e *GiftUnclaimedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// VirtualBankAccountAddedEvent represents a virtual_bank_account_added webhook event
type VirtualBankAccountAddedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *VirtualBankAccountAddedContent `json:"content"`
}

func (e *VirtualBankAccountAddedEvent) GetEventType() string     { return e.EventType }
func (e *VirtualBankAccountAddedEvent) GetEventId() string       { return e.Id }
func (e *VirtualBankAccountAddedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentIntentCreatedEvent represents a payment_intent_created webhook event
type PaymentIntentCreatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentIntentCreatedContent `json:"content"`
}

func (e *PaymentIntentCreatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentIntentCreatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentIntentCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CreditNoteCreatedWithBackdatingEvent represents a credit_note_created_with_backdating webhook event
type CreditNoteCreatedWithBackdatingEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt time.Time                               `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *CreditNoteCreatedWithBackdatingContent `json:"content"`
}

func (e *CreditNoteCreatedWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *CreditNoteCreatedWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *CreditNoteCreatedWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ContractTermTerminatedEvent represents a contract_term_terminated webhook event
type ContractTermTerminatedEvent struct {
	Id         string                         `json:"id"`
	OccurredAt time.Time                      `json:"occurred_at"`
	EventType  string                         `json:"event_type"`
	Content    *ContractTermTerminatedContent `json:"content"`
}

func (e *ContractTermTerminatedEvent) GetEventType() string     { return e.EventType }
func (e *ContractTermTerminatedEvent) GetEventId() string       { return e.Id }
func (e *ContractTermTerminatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemFamilyUpdatedEvent represents a item_family_updated webhook event
type ItemFamilyUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyUpdatedContent `json:"content"`
}

func (e *ItemFamilyUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemFamilyUpdatedEvent) GetEventId() string       { return e.Id }
func (e *ItemFamilyUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderCreatedEvent represents a order_created webhook event
type OrderCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderCreatedContent `json:"content"`
}

func (e *OrderCreatedEvent) GetEventType() string     { return e.EventType }
func (e *OrderCreatedEvent) GetEventId() string       { return e.Id }
func (e *OrderCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PriceVariantDeletedEvent represents a price_variant_deleted webhook event
type PriceVariantDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantDeletedContent `json:"content"`
}

func (e *PriceVariantDeletedEvent) GetEventType() string     { return e.EventType }
func (e *PriceVariantDeletedEvent) GetEventId() string       { return e.Id }
func (e *PriceVariantDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionMovementFailedEvent represents a subscription_movement_failed webhook event
type SubscriptionMovementFailedEvent struct {
	Id         string                             `json:"id"`
	OccurredAt time.Time                          `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *SubscriptionMovementFailedContent `json:"content"`
}

func (e *SubscriptionMovementFailedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionMovementFailedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionMovementFailedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerMovedInEvent represents a customer_moved_in webhook event
type CustomerMovedInEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerMovedInContent `json:"content"`
}

func (e *CustomerMovedInEvent) GetEventType() string     { return e.EventType }
func (e *CustomerMovedInEvent) GetEventId() string       { return e.Id }
func (e *CustomerMovedInEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionAdvanceInvoiceScheduleUpdatedEvent represents a subscription_advance_invoice_schedule_updated webhook event
type SubscriptionAdvanceInvoiceScheduleUpdatedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt time.Time                                         `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleUpdatedContent `json:"content"`
}

func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionAdvanceInvoiceScheduleUpdatedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// ItemDeletedEvent represents a item_deleted webhook event
type ItemDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemDeletedContent `json:"content"`
}

func (e *ItemDeletedEvent) GetEventType() string     { return e.EventType }
func (e *ItemDeletedEvent) GetEventId() string       { return e.Id }
func (e *ItemDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRampDraftedEvent represents a subscription_ramp_drafted webhook event
type SubscriptionRampDraftedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampDraftedContent `json:"content"`
}

func (e *SubscriptionRampDraftedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRampDraftedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRampDraftedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// DunningUpdatedEvent represents a dunning_updated webhook event
type DunningUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *DunningUpdatedContent `json:"content"`
}

func (e *DunningUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *DunningUpdatedEvent) GetEventId() string       { return e.Id }
func (e *DunningUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemEntitlementsUpdatedEvent represents a item_entitlements_updated webhook event
type ItemEntitlementsUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *ItemEntitlementsUpdatedContent `json:"content"`
}

func (e *ItemEntitlementsUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemEntitlementsUpdatedEvent) GetEventId() string       { return e.Id }
func (e *ItemEntitlementsUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TokenConsumedEvent represents a token_consumed webhook event
type TokenConsumedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *TokenConsumedContent `json:"content"`
}

func (e *TokenConsumedEvent) GetEventType() string     { return e.EventType }
func (e *TokenConsumedEvent) GetEventId() string       { return e.Id }
func (e *TokenConsumedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// HierarchyDeletedEvent represents a hierarchy_deleted webhook event
type HierarchyDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *HierarchyDeletedContent `json:"content"`
}

func (e *HierarchyDeletedEvent) GetEventType() string     { return e.EventType }
func (e *HierarchyDeletedEvent) GetEventId() string       { return e.Id }
func (e *HierarchyDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionCancellationScheduledEvent represents a subscription_cancellation_scheduled webhook event
type SubscriptionCancellationScheduledEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt time.Time                                 `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionCancellationScheduledContent `json:"content"`
}

func (e *SubscriptionCancellationScheduledEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionCancellationScheduledEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionCancellationScheduledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRenewedEvent represents a subscription_renewed webhook event
type SubscriptionRenewedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionRenewedContent `json:"content"`
}

func (e *SubscriptionRenewedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRenewedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRenewedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// FeatureUpdatedEvent represents a feature_updated webhook event
type FeatureUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureUpdatedContent `json:"content"`
}

func (e *FeatureUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *FeatureUpdatedEvent) GetEventId() string       { return e.Id }
func (e *FeatureUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// FeatureDeletedEvent represents a feature_deleted webhook event
type FeatureDeletedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureDeletedContent `json:"content"`
}

func (e *FeatureDeletedEvent) GetEventType() string     { return e.EventType }
func (e *FeatureDeletedEvent) GetEventId() string       { return e.Id }
func (e *FeatureDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemFamilyCreatedEvent represents a item_family_created webhook event
type ItemFamilyCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyCreatedContent `json:"content"`
}

func (e *ItemFamilyCreatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemFamilyCreatedEvent) GetEventId() string       { return e.Id }
func (e *ItemFamilyCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemScheduledChangeRemovedEvent represents a omnichannel_subscription_item_scheduled_change_removed webhook event
type OmnichannelSubscriptionItemScheduledChangeRemovedEvent struct {
	Id         string                                                    `json:"id"`
	OccurredAt time.Time                                                 `json:"occurred_at"`
	EventType  string                                                    `json:"event_type"`
	Content    *OmnichannelSubscriptionItemScheduledChangeRemovedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemScheduledChangeRemovedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// OmnichannelSubscriptionItemResumedEvent represents a omnichannel_subscription_item_resumed webhook event
type OmnichannelSubscriptionItemResumedEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt time.Time                                  `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemResumedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemResumedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemResumedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemResumedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PurchaseCreatedEvent represents a purchase_created webhook event
type PurchaseCreatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *PurchaseCreatedContent `json:"content"`
}

func (e *PurchaseCreatedEvent) GetEventType() string     { return e.EventType }
func (e *PurchaseCreatedEvent) GetEventId() string       { return e.Id }
func (e *PurchaseCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// EntitlementOverridesUpdatedEvent represents a entitlement_overrides_updated webhook event
type EntitlementOverridesUpdatedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt time.Time                           `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *EntitlementOverridesUpdatedContent `json:"content"`
}

func (e *EntitlementOverridesUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *EntitlementOverridesUpdatedEvent) GetEventId() string       { return e.Id }
func (e *EntitlementOverridesUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemFamilyDeletedEvent represents a item_family_deleted webhook event
type ItemFamilyDeletedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *ItemFamilyDeletedContent `json:"content"`
}

func (e *ItemFamilyDeletedEvent) GetEventType() string     { return e.EventType }
func (e *ItemFamilyDeletedEvent) GetEventId() string       { return e.Id }
func (e *ItemFamilyDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionResumptionScheduledEvent represents a subscription_resumption_scheduled webhook event
type SubscriptionResumptionScheduledEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt time.Time                               `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionResumptionScheduledContent `json:"content"`
}

func (e *SubscriptionResumptionScheduledEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionResumptionScheduledEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionResumptionScheduledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// FeatureReactivatedEvent represents a feature_reactivated webhook event
type FeatureReactivatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *FeatureReactivatedContent `json:"content"`
}

func (e *FeatureReactivatedEvent) GetEventType() string     { return e.EventType }
func (e *FeatureReactivatedEvent) GetEventId() string       { return e.Id }
func (e *FeatureReactivatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponCodesDeletedEvent represents a coupon_codes_deleted webhook event
type CouponCodesDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CouponCodesDeletedContent `json:"content"`
}

func (e *CouponCodesDeletedEvent) GetEventType() string     { return e.EventType }
func (e *CouponCodesDeletedEvent) GetEventId() string       { return e.Id }
func (e *CouponCodesDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CardExpiredEvent represents a card_expired webhook event
type CardExpiredEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardExpiredContent `json:"content"`
}

func (e *CardExpiredEvent) GetEventType() string     { return e.EventType }
func (e *CardExpiredEvent) GetEventId() string       { return e.Id }
func (e *CardExpiredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CreditNoteUpdatedEvent represents a credit_note_updated webhook event
type CreditNoteUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteUpdatedContent `json:"content"`
}

func (e *CreditNoteUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *CreditNoteUpdatedEvent) GetEventId() string       { return e.Id }
func (e *CreditNoteUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemDowngradedEvent represents a omnichannel_subscription_item_downgraded webhook event
type OmnichannelSubscriptionItemDowngradedEvent struct {
	Id         string                                        `json:"id"`
	OccurredAt time.Time                                     `json:"occurred_at"`
	EventType  string                                        `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDowngradedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemDowngradedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemDowngradedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemDowngradedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PriceVariantUpdatedEvent represents a price_variant_updated webhook event
type PriceVariantUpdatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantUpdatedContent `json:"content"`
}

func (e *PriceVariantUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *PriceVariantUpdatedEvent) GetEventId() string       { return e.Id }
func (e *PriceVariantUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PromotionalCreditsDeductedEvent represents a promotional_credits_deducted webhook event
type PromotionalCreditsDeductedEvent struct {
	Id         string                             `json:"id"`
	OccurredAt time.Time                          `json:"occurred_at"`
	EventType  string                             `json:"event_type"`
	Content    *PromotionalCreditsDeductedContent `json:"content"`
}

func (e *PromotionalCreditsDeductedEvent) GetEventType() string     { return e.EventType }
func (e *PromotionalCreditsDeductedEvent) GetEventId() string       { return e.Id }
func (e *PromotionalCreditsDeductedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRampAppliedEvent represents a subscription_ramp_applied webhook event
type SubscriptionRampAppliedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampAppliedContent `json:"content"`
}

func (e *SubscriptionRampAppliedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRampAppliedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRampAppliedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionPausedEvent represents a subscription_paused webhook event
type SubscriptionPausedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *SubscriptionPausedContent `json:"content"`
}

func (e *SubscriptionPausedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionPausedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionPausedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderReadyToProcessEvent represents a order_ready_to_process webhook event
type OrderReadyToProcessEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *OrderReadyToProcessContent `json:"content"`
}

func (e *OrderReadyToProcessEvent) GetEventType() string     { return e.EventType }
func (e *OrderReadyToProcessEvent) GetEventId() string       { return e.Id }
func (e *OrderReadyToProcessEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// FeatureCreatedEvent represents a feature_created webhook event
type FeatureCreatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *FeatureCreatedContent `json:"content"`
}

func (e *FeatureCreatedEvent) GetEventType() string     { return e.EventType }
func (e *FeatureCreatedEvent) GetEventId() string       { return e.Id }
func (e *FeatureCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TransactionDeletedEvent represents a transaction_deleted webhook event
type TransactionDeletedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionDeletedContent `json:"content"`
}

func (e *TransactionDeletedEvent) GetEventType() string     { return e.EventType }
func (e *TransactionDeletedEvent) GetEventId() string       { return e.Id }
func (e *TransactionDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CreditNoteCreatedEvent represents a credit_note_created webhook event
type CreditNoteCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteCreatedContent `json:"content"`
}

func (e *CreditNoteCreatedEvent) GetEventType() string     { return e.EventType }
func (e *CreditNoteCreatedEvent) GetEventId() string       { return e.Id }
func (e *CreditNoteCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemResubscribedEvent represents a omnichannel_subscription_item_resubscribed webhook event
type OmnichannelSubscriptionItemResubscribedEvent struct {
	Id         string                                          `json:"id"`
	OccurredAt time.Time                                       `json:"occurred_at"`
	EventType  string                                          `json:"event_type"`
	Content    *OmnichannelSubscriptionItemResubscribedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemResubscribedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemResubscribedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemResubscribedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// RecordPurchaseFailedEvent represents a record_purchase_failed webhook event
type RecordPurchaseFailedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *RecordPurchaseFailedContent `json:"content"`
}

func (e *RecordPurchaseFailedEvent) GetEventType() string     { return e.EventType }
func (e *RecordPurchaseFailedEvent) GetEventId() string       { return e.Id }
func (e *RecordPurchaseFailedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemCreatedEvent represents a item_created webhook event
type ItemCreatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemCreatedContent `json:"content"`
}

func (e *ItemCreatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemCreatedEvent) GetEventId() string       { return e.Id }
func (e *ItemCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TransactionUpdatedEvent represents a transaction_updated webhook event
type TransactionUpdatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionUpdatedContent `json:"content"`
}

func (e *TransactionUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *TransactionUpdatedEvent) GetEventId() string       { return e.Id }
func (e *TransactionUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// MrrUpdatedEvent represents a mrr_updated webhook event
type MrrUpdatedEvent struct {
	Id         string             `json:"id"`
	OccurredAt time.Time          `json:"occurred_at"`
	EventType  string             `json:"event_type"`
	Content    *MrrUpdatedContent `json:"content"`
}

func (e *MrrUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *MrrUpdatedEvent) GetEventId() string       { return e.Id }
func (e *MrrUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// UnbilledChargesInvoicedEvent represents a unbilled_charges_invoiced webhook event
type UnbilledChargesInvoicedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *UnbilledChargesInvoicedContent `json:"content"`
}

func (e *UnbilledChargesInvoicedEvent) GetEventType() string     { return e.EventType }
func (e *UnbilledChargesInvoicedEvent) GetEventId() string       { return e.Id }
func (e *UnbilledChargesInvoicedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemPriceUpdatedEvent represents a item_price_updated webhook event
type ItemPriceUpdatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceUpdatedContent `json:"content"`
}

func (e *ItemPriceUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemPriceUpdatedEvent) GetEventId() string       { return e.Id }
func (e *ItemPriceUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponCodesUpdatedEvent represents a coupon_codes_updated webhook event
type CouponCodesUpdatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CouponCodesUpdatedContent `json:"content"`
}

func (e *CouponCodesUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *CouponCodesUpdatedEvent) GetEventId() string       { return e.Id }
func (e *CouponCodesUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// VirtualBankAccountUpdatedEvent represents a virtual_bank_account_updated webhook event
type VirtualBankAccountUpdatedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt time.Time                         `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *VirtualBankAccountUpdatedContent `json:"content"`
}

func (e *VirtualBankAccountUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *VirtualBankAccountUpdatedEvent) GetEventId() string       { return e.Id }
func (e *VirtualBankAccountUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ContractTermCreatedEvent represents a contract_term_created webhook event
type ContractTermCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *ContractTermCreatedContent `json:"content"`
}

func (e *ContractTermCreatedEvent) GetEventType() string     { return e.EventType }
func (e *ContractTermCreatedEvent) GetEventId() string       { return e.Id }
func (e *ContractTermCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionChangedEvent represents a subscription_changed webhook event
type SubscriptionChangedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionChangedContent `json:"content"`
}

func (e *SubscriptionChangedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionChangedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionChangedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentFailedEvent represents a payment_failed webhook event
type PaymentFailedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *PaymentFailedContent `json:"content"`
}

func (e *PaymentFailedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentFailedEvent) GetEventId() string       { return e.Id }
func (e *PaymentFailedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CreditNoteDeletedEvent represents a credit_note_deleted webhook event
type CreditNoteDeletedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *CreditNoteDeletedContent `json:"content"`
}

func (e *CreditNoteDeletedEvent) GetEventType() string     { return e.EventType }
func (e *CreditNoteDeletedEvent) GetEventId() string       { return e.Id }
func (e *CreditNoteDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TaxWithheldRefundedEvent represents a tax_withheld_refunded webhook event
type TaxWithheldRefundedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *TaxWithheldRefundedContent `json:"content"`
}

func (e *TaxWithheldRefundedEvent) GetEventType() string     { return e.EventType }
func (e *TaxWithheldRefundedEvent) GetEventId() string       { return e.Id }
func (e *TaxWithheldRefundedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ContractTermCompletedEvent represents a contract_term_completed webhook event
type ContractTermCompletedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *ContractTermCompletedContent `json:"content"`
}

func (e *ContractTermCompletedEvent) GetEventType() string     { return e.EventType }
func (e *ContractTermCompletedEvent) GetEventId() string       { return e.Id }
func (e *ContractTermCompletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSchedulesUpdatedEvent represents a payment_schedules_updated webhook event
type PaymentSchedulesUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PaymentSchedulesUpdatedContent `json:"content"`
}

func (e *PaymentSchedulesUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSchedulesUpdatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentSchedulesUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemExpiredEvent represents a omnichannel_subscription_item_expired webhook event
type OmnichannelSubscriptionItemExpiredEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt time.Time                                  `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *OmnichannelSubscriptionItemExpiredContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemExpiredEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemExpiredEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemExpiredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CardUpdatedEvent represents a card_updated webhook event
type CardUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardUpdatedContent `json:"content"`
}

func (e *CardUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *CardUpdatedEvent) GetEventId() string       { return e.Id }
func (e *CardUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerCreatedEvent represents a customer_created webhook event
type CustomerCreatedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerCreatedContent `json:"content"`
}

func (e *CustomerCreatedEvent) GetEventType() string     { return e.EventType }
func (e *CustomerCreatedEvent) GetEventId() string       { return e.Id }
func (e *CustomerCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRenewalReminderEvent represents a subscription_renewal_reminder webhook event
type SubscriptionRenewalReminderEvent struct {
	Id         string                              `json:"id"`
	OccurredAt time.Time                           `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *SubscriptionRenewalReminderContent `json:"content"`
}

func (e *SubscriptionRenewalReminderEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRenewalReminderEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRenewalReminderEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderDeliveredEvent represents a order_delivered webhook event
type OrderDeliveredEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *OrderDeliveredContent `json:"content"`
}

func (e *OrderDeliveredEvent) GetEventType() string     { return e.EventType }
func (e *OrderDeliveredEvent) GetEventId() string       { return e.Id }
func (e *OrderDeliveredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemCancellationScheduledEvent represents a omnichannel_subscription_item_cancellation_scheduled webhook event
type OmnichannelSubscriptionItemCancellationScheduledEvent struct {
	Id         string                                                   `json:"id"`
	OccurredAt time.Time                                                `json:"occurred_at"`
	EventType  string                                                   `json:"event_type"`
	Content    *OmnichannelSubscriptionItemCancellationScheduledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemCancellationScheduledEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// OmnichannelSubscriptionItemGracePeriodExpiredEvent represents a omnichannel_subscription_item_grace_period_expired webhook event
type OmnichannelSubscriptionItemGracePeriodExpiredEvent struct {
	Id         string                                                `json:"id"`
	OccurredAt time.Time                                             `json:"occurred_at"`
	EventType  string                                                `json:"event_type"`
	Content    *OmnichannelSubscriptionItemGracePeriodExpiredContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemGracePeriodExpiredEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// CouponCodesAddedEvent represents a coupon_codes_added webhook event
type CouponCodesAddedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponCodesAddedContent `json:"content"`
}

func (e *CouponCodesAddedEvent) GetEventType() string     { return e.EventType }
func (e *CouponCodesAddedEvent) GetEventId() string       { return e.Id }
func (e *CouponCodesAddedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// GiftCancelledEvent represents a gift_cancelled webhook event
type GiftCancelledEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *GiftCancelledContent `json:"content"`
}

func (e *GiftCancelledEvent) GetEventType() string     { return e.EventType }
func (e *GiftCancelledEvent) GetEventId() string       { return e.Id }
func (e *GiftCancelledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderCancelledEvent represents a order_cancelled webhook event
type OrderCancelledEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *OrderCancelledContent `json:"content"`
}

func (e *OrderCancelledEvent) GetEventType() string     { return e.EventType }
func (e *OrderCancelledEvent) GetEventId() string       { return e.Id }
func (e *OrderCancelledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponDeletedEvent represents a coupon_deleted webhook event
type CouponDeletedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponDeletedContent `json:"content"`
}

func (e *CouponDeletedEvent) GetEventType() string     { return e.EventType }
func (e *CouponDeletedEvent) GetEventId() string       { return e.Id }
func (e *CouponDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionScheduledChangesRemovedEvent represents a subscription_scheduled_changes_removed webhook event
type SubscriptionScheduledChangesRemovedEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt time.Time                                   `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *SubscriptionScheduledChangesRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledChangesRemovedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionScheduledChangesRemovedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionScheduledChangesRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PendingInvoiceCreatedEvent represents a pending_invoice_created webhook event
type PendingInvoiceCreatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PendingInvoiceCreatedContent `json:"content"`
}

func (e *PendingInvoiceCreatedEvent) GetEventType() string     { return e.EventType }
func (e *PendingInvoiceCreatedEvent) GetEventId() string       { return e.Id }
func (e *PendingInvoiceCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// EntitlementOverridesAutoRemovedEvent represents a entitlement_overrides_auto_removed webhook event
type EntitlementOverridesAutoRemovedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt time.Time                               `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *EntitlementOverridesAutoRemovedContent `json:"content"`
}

func (e *EntitlementOverridesAutoRemovedEvent) GetEventType() string     { return e.EventType }
func (e *EntitlementOverridesAutoRemovedEvent) GetEventId() string       { return e.Id }
func (e *EntitlementOverridesAutoRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemUpgradedEvent represents a omnichannel_subscription_item_upgraded webhook event
type OmnichannelSubscriptionItemUpgradedEvent struct {
	Id         string                                      `json:"id"`
	OccurredAt time.Time                                   `json:"occurred_at"`
	EventType  string                                      `json:"event_type"`
	Content    *OmnichannelSubscriptionItemUpgradedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemUpgradedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemUpgradedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemUpgradedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionBusinessEntityChangedEvent represents a subscription_business_entity_changed webhook event
type SubscriptionBusinessEntityChangedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt time.Time                                 `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionBusinessEntityChangedContent `json:"content"`
}

func (e *SubscriptionBusinessEntityChangedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionBusinessEntityChangedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionBusinessEntityChangedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelOneTimeOrderCreatedEvent represents a omnichannel_one_time_order_created webhook event
type OmnichannelOneTimeOrderCreatedEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt time.Time                              `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelOneTimeOrderCreatedContent `json:"content"`
}

func (e *OmnichannelOneTimeOrderCreatedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelOneTimeOrderCreatedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelOneTimeOrderCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSourceDeletedEvent represents a payment_source_deleted webhook event
type PaymentSourceDeletedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceDeletedContent `json:"content"`
}

func (e *PaymentSourceDeletedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSourceDeletedEvent) GetEventId() string       { return e.Id }
func (e *PaymentSourceDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemCancelledEvent represents a omnichannel_subscription_item_cancelled webhook event
type OmnichannelSubscriptionItemCancelledEvent struct {
	Id         string                                       `json:"id"`
	OccurredAt time.Time                                    `json:"occurred_at"`
	EventType  string                                       `json:"event_type"`
	Content    *OmnichannelSubscriptionItemCancelledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemCancelledEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionItemCancelledEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionItemCancelledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// QuoteDeletedEvent represents a quote_deleted webhook event
type QuoteDeletedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteDeletedContent `json:"content"`
}

func (e *QuoteDeletedEvent) GetEventType() string     { return e.EventType }
func (e *QuoteDeletedEvent) GetEventId() string       { return e.Id }
func (e *QuoteDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// InvoiceUpdatedEvent represents a invoice_updated webhook event
type InvoiceUpdatedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *InvoiceUpdatedContent `json:"content"`
}

func (e *InvoiceUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *InvoiceUpdatedEvent) GetEventId() string       { return e.Id }
func (e *InvoiceUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionAdvanceInvoiceScheduleRemovedEvent represents a subscription_advance_invoice_schedule_removed webhook event
type SubscriptionAdvanceInvoiceScheduleRemovedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt time.Time                                         `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *SubscriptionAdvanceInvoiceScheduleRemovedContent `json:"content"`
}

func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionAdvanceInvoiceScheduleRemovedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// CardDeletedEvent represents a card_deleted webhook event
type CardDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *CardDeletedContent `json:"content"`
}

func (e *CardDeletedEvent) GetEventType() string     { return e.EventType }
func (e *CardDeletedEvent) GetEventId() string       { return e.Id }
func (e *CardDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderReadyToShipEvent represents a order_ready_to_ship webhook event
type OrderReadyToShipEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *OrderReadyToShipContent `json:"content"`
}

func (e *OrderReadyToShipEvent) GetEventType() string     { return e.EventType }
func (e *OrderReadyToShipEvent) GetEventId() string       { return e.Id }
func (e *OrderReadyToShipEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionMovedOutEvent represents a subscription_moved_out webhook event
type SubscriptionMovedOutEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *SubscriptionMovedOutContent `json:"content"`
}

func (e *SubscriptionMovedOutEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionMovedOutEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionMovedOutEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentScheduleSchemeCreatedEvent represents a payment_schedule_scheme_created webhook event
type PaymentScheduleSchemeCreatedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt time.Time                            `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *PaymentScheduleSchemeCreatedContent `json:"content"`
}

func (e *PaymentScheduleSchemeCreatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentScheduleSchemeCreatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentScheduleSchemeCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// BusinessEntityUpdatedEvent represents a business_entity_updated webhook event
type BusinessEntityUpdatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityUpdatedContent `json:"content"`
}

func (e *BusinessEntityUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *BusinessEntityUpdatedEvent) GetEventId() string       { return e.Id }
func (e *BusinessEntityUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionScheduledResumptionRemovedEvent represents a subscription_scheduled_resumption_removed webhook event
type SubscriptionScheduledResumptionRemovedEvent struct {
	Id         string                                         `json:"id"`
	OccurredAt time.Time                                      `json:"occurred_at"`
	EventType  string                                         `json:"event_type"`
	Content    *SubscriptionScheduledResumptionRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledResumptionRemovedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionScheduledResumptionRemovedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionScheduledResumptionRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentInitiatedEvent represents a payment_initiated webhook event
type PaymentInitiatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *PaymentInitiatedContent `json:"content"`
}

func (e *PaymentInitiatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentInitiatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentInitiatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// FeatureArchivedEvent represents a feature_archived webhook event
type FeatureArchivedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *FeatureArchivedContent `json:"content"`
}

func (e *FeatureArchivedEvent) GetEventType() string     { return e.EventType }
func (e *FeatureArchivedEvent) GetEventId() string       { return e.Id }
func (e *FeatureArchivedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionReactivatedWithBackdatingEvent represents a subscription_reactivated_with_backdating webhook event
type SubscriptionReactivatedWithBackdatingEvent struct {
	Id         string                                        `json:"id"`
	OccurredAt time.Time                                     `json:"occurred_at"`
	EventType  string                                        `json:"event_type"`
	Content    *SubscriptionReactivatedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionReactivatedWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionReactivatedWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionReactivatedWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionImportedEvent represents a omnichannel_subscription_imported webhook event
type OmnichannelSubscriptionImportedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt time.Time                               `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *OmnichannelSubscriptionImportedContent `json:"content"`
}

func (e *OmnichannelSubscriptionImportedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionImportedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionImportedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TokenExpiredEvent represents a token_expired webhook event
type TokenExpiredEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *TokenExpiredContent `json:"content"`
}

func (e *TokenExpiredEvent) GetEventType() string     { return e.EventType }
func (e *TokenExpiredEvent) GetEventId() string       { return e.Id }
func (e *TokenExpiredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CardAddedEvent represents a card_added webhook event
type CardAddedEvent struct {
	Id         string            `json:"id"`
	OccurredAt time.Time         `json:"occurred_at"`
	EventType  string            `json:"event_type"`
	Content    *CardAddedContent `json:"content"`
}

func (e *CardAddedEvent) GetEventType() string     { return e.EventType }
func (e *CardAddedEvent) GetEventId() string       { return e.Id }
func (e *CardAddedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponCreatedEvent represents a coupon_created webhook event
type CouponCreatedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *CouponCreatedContent `json:"content"`
}

func (e *CouponCreatedEvent) GetEventType() string     { return e.EventType }
func (e *CouponCreatedEvent) GetEventId() string       { return e.Id }
func (e *CouponCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// RuleDeletedEvent represents a rule_deleted webhook event
type RuleDeletedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleDeletedContent `json:"content"`
}

func (e *RuleDeletedEvent) GetEventType() string     { return e.EventType }
func (e *RuleDeletedEvent) GetEventId() string       { return e.Id }
func (e *RuleDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemPriceEntitlementsUpdatedEvent represents a item_price_entitlements_updated webhook event
type ItemPriceEntitlementsUpdatedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt time.Time                            `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *ItemPriceEntitlementsUpdatedContent `json:"content"`
}

func (e *ItemPriceEntitlementsUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemPriceEntitlementsUpdatedEvent) GetEventId() string       { return e.Id }
func (e *ItemPriceEntitlementsUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemPriceDeletedEvent represents a item_price_deleted webhook event
type ItemPriceDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceDeletedContent `json:"content"`
}

func (e *ItemPriceDeletedEvent) GetEventType() string     { return e.EventType }
func (e *ItemPriceDeletedEvent) GetEventId() string       { return e.Id }
func (e *ItemPriceDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// VirtualBankAccountDeletedEvent represents a virtual_bank_account_deleted webhook event
type VirtualBankAccountDeletedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt time.Time                         `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *VirtualBankAccountDeletedContent `json:"content"`
}

func (e *VirtualBankAccountDeletedEvent) GetEventType() string     { return e.EventType }
func (e *VirtualBankAccountDeletedEvent) GetEventId() string       { return e.Id }
func (e *VirtualBankAccountDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentScheduleSchemeDeletedEvent represents a payment_schedule_scheme_deleted webhook event
type PaymentScheduleSchemeDeletedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt time.Time                            `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *PaymentScheduleSchemeDeletedContent `json:"content"`
}

func (e *PaymentScheduleSchemeDeletedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentScheduleSchemeDeletedEvent) GetEventId() string       { return e.Id }
func (e *PaymentScheduleSchemeDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionCreatedEvent represents a subscription_created webhook event
type SubscriptionCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionCreatedContent `json:"content"`
}

func (e *SubscriptionCreatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionCreatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionEntitlementsCreatedEvent represents a subscription_entitlements_created webhook event
type SubscriptionEntitlementsCreatedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt time.Time                               `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionEntitlementsCreatedContent `json:"content"`
}

func (e *SubscriptionEntitlementsCreatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionEntitlementsCreatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionEntitlementsCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderReturnedEvent represents a order_returned webhook event
type OrderReturnedEvent struct {
	Id         string                `json:"id"`
	OccurredAt time.Time             `json:"occurred_at"`
	EventType  string                `json:"event_type"`
	Content    *OrderReturnedContent `json:"content"`
}

func (e *OrderReturnedEvent) GetEventType() string     { return e.EventType }
func (e *OrderReturnedEvent) GetEventId() string       { return e.Id }
func (e *OrderReturnedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionDeletedEvent represents a subscription_deleted webhook event
type SubscriptionDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionDeletedContent `json:"content"`
}

func (e *SubscriptionDeletedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionDeletedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSourceAddedEvent represents a payment_source_added webhook event
type PaymentSourceAddedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *PaymentSourceAddedContent `json:"content"`
}

func (e *PaymentSourceAddedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSourceAddedEvent) GetEventId() string       { return e.Id }
func (e *PaymentSourceAddedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionMovedInEvent represents a subscription_moved_in webhook event
type SubscriptionMovedInEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionMovedInContent `json:"content"`
}

func (e *SubscriptionMovedInEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionMovedInEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionMovedInEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemPriceCreatedEvent represents a item_price_created webhook event
type ItemPriceCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *ItemPriceCreatedContent `json:"content"`
}

func (e *ItemPriceCreatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemPriceCreatedEvent) GetEventId() string       { return e.Id }
func (e *ItemPriceCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionScheduledCancellationRemovedEvent represents a subscription_scheduled_cancellation_removed webhook event
type SubscriptionScheduledCancellationRemovedEvent struct {
	Id         string                                           `json:"id"`
	OccurredAt time.Time                                        `json:"occurred_at"`
	EventType  string                                           `json:"event_type"`
	Content    *SubscriptionScheduledCancellationRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledCancellationRemovedEvent) GetEventType() string { return e.EventType }
func (e *SubscriptionScheduledCancellationRemovedEvent) GetEventId() string   { return e.Id }
func (e *SubscriptionScheduledCancellationRemovedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// PaymentRefundedEvent represents a payment_refunded webhook event
type PaymentRefundedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *PaymentRefundedContent `json:"content"`
}

func (e *PaymentRefundedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentRefundedEvent) GetEventId() string       { return e.Id }
func (e *PaymentRefundedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// UsageFileIngestedEvent represents a usage_file_ingested webhook event
type UsageFileIngestedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *UsageFileIngestedContent `json:"content"`
}

func (e *UsageFileIngestedEvent) GetEventType() string     { return e.EventType }
func (e *UsageFileIngestedEvent) GetEventId() string       { return e.Id }
func (e *UsageFileIngestedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionMovedInEvent represents a omnichannel_subscription_moved_in webhook event
type OmnichannelSubscriptionMovedInEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt time.Time                              `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelSubscriptionMovedInContent `json:"content"`
}

func (e *OmnichannelSubscriptionMovedInEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionMovedInEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionMovedInEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// DifferentialPriceCreatedEvent represents a differential_price_created webhook event
type DifferentialPriceCreatedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt time.Time                        `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceCreatedContent `json:"content"`
}

func (e *DifferentialPriceCreatedEvent) GetEventType() string     { return e.EventType }
func (e *DifferentialPriceCreatedEvent) GetEventId() string       { return e.Id }
func (e *DifferentialPriceCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TransactionCreatedEvent represents a transaction_created webhook event
type TransactionCreatedEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *TransactionCreatedContent `json:"content"`
}

func (e *TransactionCreatedEvent) GetEventType() string     { return e.EventType }
func (e *TransactionCreatedEvent) GetEventId() string       { return e.Id }
func (e *TransactionCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSucceededEvent represents a payment_succeeded webhook event
type PaymentSucceededEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *PaymentSucceededContent `json:"content"`
}

func (e *PaymentSucceededEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSucceededEvent) GetEventId() string       { return e.Id }
func (e *PaymentSucceededEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionCanceledWithBackdatingEvent represents a subscription_canceled_with_backdating webhook event
type SubscriptionCanceledWithBackdatingEvent struct {
	Id         string                                     `json:"id"`
	OccurredAt time.Time                                  `json:"occurred_at"`
	EventType  string                                     `json:"event_type"`
	Content    *SubscriptionCanceledWithBackdatingContent `json:"content"`
}

func (e *SubscriptionCanceledWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionCanceledWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionCanceledWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// UnbilledChargesVoidedEvent represents a unbilled_charges_voided webhook event
type UnbilledChargesVoidedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *UnbilledChargesVoidedContent `json:"content"`
}

func (e *UnbilledChargesVoidedEvent) GetEventType() string     { return e.EventType }
func (e *UnbilledChargesVoidedEvent) GetEventId() string       { return e.Id }
func (e *UnbilledChargesVoidedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// QuoteCreatedEvent represents a quote_created webhook event
type QuoteCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteCreatedContent `json:"content"`
}

func (e *QuoteCreatedEvent) GetEventType() string     { return e.EventType }
func (e *QuoteCreatedEvent) GetEventId() string       { return e.Id }
func (e *QuoteCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponSetDeletedEvent represents a coupon_set_deleted webhook event
type CouponSetDeletedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetDeletedContent `json:"content"`
}

func (e *CouponSetDeletedEvent) GetEventType() string     { return e.EventType }
func (e *CouponSetDeletedEvent) GetEventId() string       { return e.Id }
func (e *CouponSetDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// AttachedItemCreatedEvent represents a attached_item_created webhook event
type AttachedItemCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemCreatedContent `json:"content"`
}

func (e *AttachedItemCreatedEvent) GetEventType() string     { return e.EventType }
func (e *AttachedItemCreatedEvent) GetEventId() string       { return e.Id }
func (e *AttachedItemCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SalesOrderCreatedEvent represents a sales_order_created webhook event
type SalesOrderCreatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *SalesOrderCreatedContent `json:"content"`
}

func (e *SalesOrderCreatedEvent) GetEventType() string     { return e.EventType }
func (e *SalesOrderCreatedEvent) GetEventId() string       { return e.Id }
func (e *SalesOrderCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerChangedEvent represents a customer_changed webhook event
type CustomerChangedEvent struct {
	Id         string                  `json:"id"`
	OccurredAt time.Time               `json:"occurred_at"`
	EventType  string                  `json:"event_type"`
	Content    *CustomerChangedContent `json:"content"`
}

func (e *CustomerChangedEvent) GetEventType() string     { return e.EventType }
func (e *CustomerChangedEvent) GetEventId() string       { return e.Id }
func (e *CustomerChangedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionStartedEvent represents a subscription_started webhook event
type SubscriptionStartedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *SubscriptionStartedContent `json:"content"`
}

func (e *SubscriptionStartedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionStartedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionStartedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionActivatedEvent represents a subscription_activated webhook event
type SubscriptionActivatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *SubscriptionActivatedContent `json:"content"`
}

func (e *SubscriptionActivatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionActivatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionActivatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSourceExpiringEvent represents a payment_source_expiring webhook event
type PaymentSourceExpiringEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PaymentSourceExpiringContent `json:"content"`
}

func (e *PaymentSourceExpiringEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSourceExpiringEvent) GetEventId() string       { return e.Id }
func (e *PaymentSourceExpiringEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionReactivatedEvent represents a subscription_reactivated webhook event
type SubscriptionReactivatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionReactivatedContent `json:"content"`
}

func (e *SubscriptionReactivatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionReactivatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionReactivatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderUpdatedEvent represents a order_updated webhook event
type OrderUpdatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderUpdatedContent `json:"content"`
}

func (e *OrderUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *OrderUpdatedEvent) GetEventId() string       { return e.Id }
func (e *OrderUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionScheduledPauseRemovedEvent represents a subscription_scheduled_pause_removed webhook event
type SubscriptionScheduledPauseRemovedEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt time.Time                                 `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionScheduledPauseRemovedContent `json:"content"`
}

func (e *SubscriptionScheduledPauseRemovedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionScheduledPauseRemovedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionScheduledPauseRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionCancellationReminderEvent represents a subscription_cancellation_reminder webhook event
type SubscriptionCancellationReminderEvent struct {
	Id         string                                   `json:"id"`
	OccurredAt time.Time                                `json:"occurred_at"`
	EventType  string                                   `json:"event_type"`
	Content    *SubscriptionCancellationReminderContent `json:"content"`
}

func (e *SubscriptionCancellationReminderEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionCancellationReminderEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionCancellationReminderEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionCreatedWithBackdatingEvent represents a subscription_created_with_backdating webhook event
type SubscriptionCreatedWithBackdatingEvent struct {
	Id         string                                    `json:"id"`
	OccurredAt time.Time                                 `json:"occurred_at"`
	EventType  string                                    `json:"event_type"`
	Content    *SubscriptionCreatedWithBackdatingContent `json:"content"`
}

func (e *SubscriptionCreatedWithBackdatingEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionCreatedWithBackdatingEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionCreatedWithBackdatingEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRampCreatedEvent represents a subscription_ramp_created webhook event
type SubscriptionRampCreatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampCreatedContent `json:"content"`
}

func (e *SubscriptionRampCreatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRampCreatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRampCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderDeletedEvent represents a order_deleted webhook event
type OrderDeletedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *OrderDeletedContent `json:"content"`
}

func (e *OrderDeletedEvent) GetEventType() string     { return e.EventType }
func (e *OrderDeletedEvent) GetEventId() string       { return e.Id }
func (e *OrderDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemPauseScheduledEvent represents a omnichannel_subscription_item_pause_scheduled webhook event
type OmnichannelSubscriptionItemPauseScheduledEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt time.Time                                         `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemPauseScheduledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemPauseScheduledEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// GiftUpdatedEvent represents a gift_updated webhook event
type GiftUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *GiftUpdatedContent `json:"content"`
}

func (e *GiftUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *GiftUpdatedEvent) GetEventId() string       { return e.Id }
func (e *GiftUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionTrialExtendedEvent represents a subscription_trial_extended webhook event
type SubscriptionTrialExtendedEvent struct {
	Id         string                            `json:"id"`
	OccurredAt time.Time                         `json:"occurred_at"`
	EventType  string                            `json:"event_type"`
	Content    *SubscriptionTrialExtendedContent `json:"content"`
}

func (e *SubscriptionTrialExtendedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionTrialExtendedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionTrialExtendedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemGracePeriodStartedEvent represents a omnichannel_subscription_item_grace_period_started webhook event
type OmnichannelSubscriptionItemGracePeriodStartedEvent struct {
	Id         string                                                `json:"id"`
	OccurredAt time.Time                                             `json:"occurred_at"`
	EventType  string                                                `json:"event_type"`
	Content    *OmnichannelSubscriptionItemGracePeriodStartedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetEventId() string { return e.Id }
func (e *OmnichannelSubscriptionItemGracePeriodStartedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// CardExpiryReminderEvent represents a card_expiry_reminder webhook event
type CardExpiryReminderEvent struct {
	Id         string                     `json:"id"`
	OccurredAt time.Time                  `json:"occurred_at"`
	EventType  string                     `json:"event_type"`
	Content    *CardExpiryReminderContent `json:"content"`
}

func (e *CardExpiryReminderEvent) GetEventType() string     { return e.EventType }
func (e *CardExpiryReminderEvent) GetEventId() string       { return e.Id }
func (e *CardExpiryReminderEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TokenCreatedEvent represents a token_created webhook event
type TokenCreatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *TokenCreatedContent `json:"content"`
}

func (e *TokenCreatedEvent) GetEventType() string     { return e.EventType }
func (e *TokenCreatedEvent) GetEventId() string       { return e.Id }
func (e *TokenCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PromotionalCreditsAddedEvent represents a promotional_credits_added webhook event
type PromotionalCreditsAddedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *PromotionalCreditsAddedContent `json:"content"`
}

func (e *PromotionalCreditsAddedEvent) GetEventType() string     { return e.EventType }
func (e *PromotionalCreditsAddedEvent) GetEventId() string       { return e.Id }
func (e *PromotionalCreditsAddedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRampUpdatedEvent represents a subscription_ramp_updated webhook event
type SubscriptionRampUpdatedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampUpdatedContent `json:"content"`
}

func (e *SubscriptionRampUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRampUpdatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRampUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerEntitlementsUpdatedEvent represents a customer_entitlements_updated webhook event
type CustomerEntitlementsUpdatedEvent struct {
	Id         string                              `json:"id"`
	OccurredAt time.Time                           `json:"occurred_at"`
	EventType  string                              `json:"event_type"`
	Content    *CustomerEntitlementsUpdatedContent `json:"content"`
}

func (e *CustomerEntitlementsUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *CustomerEntitlementsUpdatedEvent) GetEventId() string       { return e.Id }
func (e *CustomerEntitlementsUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSourceExpiredEvent represents a payment_source_expired webhook event
type PaymentSourceExpiredEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceExpiredContent `json:"content"`
}

func (e *PaymentSourceExpiredEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSourceExpiredEvent) GetEventId() string       { return e.Id }
func (e *PaymentSourceExpiredEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CustomerMovedOutEvent represents a customer_moved_out webhook event
type CustomerMovedOutEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CustomerMovedOutContent `json:"content"`
}

func (e *CustomerMovedOutEvent) GetEventType() string     { return e.EventType }
func (e *CustomerMovedOutEvent) GetEventId() string       { return e.Id }
func (e *CustomerMovedOutEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionEntitlementsUpdatedEvent represents a subscription_entitlements_updated webhook event
type SubscriptionEntitlementsUpdatedEvent struct {
	Id         string                                  `json:"id"`
	OccurredAt time.Time                               `json:"occurred_at"`
	EventType  string                                  `json:"event_type"`
	Content    *SubscriptionEntitlementsUpdatedContent `json:"content"`
}

func (e *SubscriptionEntitlementsUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionEntitlementsUpdatedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionEntitlementsUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemDunningExpiredEvent represents a omnichannel_subscription_item_dunning_expired webhook event
type OmnichannelSubscriptionItemDunningExpiredEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt time.Time                                         `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDunningExpiredContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemDunningExpiredEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// HierarchyCreatedEvent represents a hierarchy_created webhook event
type HierarchyCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *HierarchyCreatedContent `json:"content"`
}

func (e *HierarchyCreatedEvent) GetEventType() string     { return e.EventType }
func (e *HierarchyCreatedEvent) GetEventId() string       { return e.Id }
func (e *HierarchyCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// AttachedItemDeletedEvent represents a attached_item_deleted webhook event
type AttachedItemDeletedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemDeletedContent `json:"content"`
}

func (e *AttachedItemDeletedEvent) GetEventType() string     { return e.EventType }
func (e *AttachedItemDeletedEvent) GetEventId() string       { return e.Id }
func (e *AttachedItemDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemScheduledCancellationRemovedEvent represents a omnichannel_subscription_item_scheduled_cancellation_removed webhook event
type OmnichannelSubscriptionItemScheduledCancellationRemovedEvent struct {
	Id         string                                                          `json:"id"`
	OccurredAt time.Time                                                       `json:"occurred_at"`
	EventType  string                                                          `json:"event_type"`
	Content    *OmnichannelSubscriptionItemScheduledCancellationRemovedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetEventType() string {
	return e.EventType
}
func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetEventId() string {
	return e.Id
}
func (e *OmnichannelSubscriptionItemScheduledCancellationRemovedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// ItemUpdatedEvent represents a item_updated webhook event
type ItemUpdatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *ItemUpdatedContent `json:"content"`
}

func (e *ItemUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *ItemUpdatedEvent) GetEventId() string       { return e.Id }
func (e *ItemUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// CouponSetCreatedEvent represents a coupon_set_created webhook event
type CouponSetCreatedEvent struct {
	Id         string                   `json:"id"`
	OccurredAt time.Time                `json:"occurred_at"`
	EventType  string                   `json:"event_type"`
	Content    *CouponSetCreatedContent `json:"content"`
}

func (e *CouponSetCreatedEvent) GetEventType() string     { return e.EventType }
func (e *CouponSetCreatedEvent) GetEventId() string       { return e.Id }
func (e *CouponSetCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentIntentUpdatedEvent represents a payment_intent_updated webhook event
type PaymentIntentUpdatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentIntentUpdatedContent `json:"content"`
}

func (e *PaymentIntentUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentIntentUpdatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentIntentUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OrderResentEvent represents a order_resent webhook event
type OrderResentEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *OrderResentContent `json:"content"`
}

func (e *OrderResentEvent) GetEventType() string     { return e.EventType }
func (e *OrderResentEvent) GetEventId() string       { return e.Id }
func (e *OrderResentEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionCreatedEvent represents a omnichannel_subscription_created webhook event
type OmnichannelSubscriptionCreatedEvent struct {
	Id         string                                 `json:"id"`
	OccurredAt time.Time                              `json:"occurred_at"`
	EventType  string                                 `json:"event_type"`
	Content    *OmnichannelSubscriptionCreatedContent `json:"content"`
}

func (e *OmnichannelSubscriptionCreatedEvent) GetEventType() string     { return e.EventType }
func (e *OmnichannelSubscriptionCreatedEvent) GetEventId() string       { return e.Id }
func (e *OmnichannelSubscriptionCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// TaxWithheldRecordedEvent represents a tax_withheld_recorded webhook event
type TaxWithheldRecordedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *TaxWithheldRecordedContent `json:"content"`
}

func (e *TaxWithheldRecordedEvent) GetEventType() string     { return e.EventType }
func (e *TaxWithheldRecordedEvent) GetEventId() string       { return e.Id }
func (e *TaxWithheldRecordedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PriceVariantCreatedEvent represents a price_variant_created webhook event
type PriceVariantCreatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *PriceVariantCreatedContent `json:"content"`
}

func (e *PriceVariantCreatedEvent) GetEventType() string     { return e.EventType }
func (e *PriceVariantCreatedEvent) GetEventId() string       { return e.Id }
func (e *PriceVariantCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// DifferentialPriceDeletedEvent represents a differential_price_deleted webhook event
type DifferentialPriceDeletedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt time.Time                        `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *DifferentialPriceDeletedContent `json:"content"`
}

func (e *DifferentialPriceDeletedEvent) GetEventType() string     { return e.EventType }
func (e *DifferentialPriceDeletedEvent) GetEventId() string       { return e.Id }
func (e *DifferentialPriceDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionItemsRenewedEvent represents a subscription_items_renewed webhook event
type SubscriptionItemsRenewedEvent struct {
	Id         string                           `json:"id"`
	OccurredAt time.Time                        `json:"occurred_at"`
	EventType  string                           `json:"event_type"`
	Content    *SubscriptionItemsRenewedContent `json:"content"`
}

func (e *SubscriptionItemsRenewedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionItemsRenewedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionItemsRenewedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// RuleCreatedEvent represents a rule_created webhook event
type RuleCreatedEvent struct {
	Id         string              `json:"id"`
	OccurredAt time.Time           `json:"occurred_at"`
	EventType  string              `json:"event_type"`
	Content    *RuleCreatedContent `json:"content"`
}

func (e *RuleCreatedEvent) GetEventType() string     { return e.EventType }
func (e *RuleCreatedEvent) GetEventId() string       { return e.Id }
func (e *RuleCreatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ContractTermCancelledEvent represents a contract_term_cancelled webhook event
type ContractTermCancelledEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *ContractTermCancelledContent `json:"content"`
}

func (e *ContractTermCancelledEvent) GetEventType() string     { return e.EventType }
func (e *ContractTermCancelledEvent) GetEventId() string       { return e.Id }
func (e *ContractTermCancelledEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ContractTermRenewedEvent represents a contract_term_renewed webhook event
type ContractTermRenewedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *ContractTermRenewedContent `json:"content"`
}

func (e *ContractTermRenewedEvent) GetEventType() string     { return e.EventType }
func (e *ContractTermRenewedEvent) GetEventId() string       { return e.Id }
func (e *ContractTermRenewedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// InvoiceDeletedEvent represents a invoice_deleted webhook event
type InvoiceDeletedEvent struct {
	Id         string                 `json:"id"`
	OccurredAt time.Time              `json:"occurred_at"`
	EventType  string                 `json:"event_type"`
	Content    *InvoiceDeletedContent `json:"content"`
}

func (e *InvoiceDeletedEvent) GetEventType() string     { return e.EventType }
func (e *InvoiceDeletedEvent) GetEventId() string       { return e.Id }
func (e *InvoiceDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ItemPriceEntitlementsRemovedEvent represents a item_price_entitlements_removed webhook event
type ItemPriceEntitlementsRemovedEvent struct {
	Id         string                               `json:"id"`
	OccurredAt time.Time                            `json:"occurred_at"`
	EventType  string                               `json:"event_type"`
	Content    *ItemPriceEntitlementsRemovedContent `json:"content"`
}

func (e *ItemPriceEntitlementsRemovedEvent) GetEventType() string     { return e.EventType }
func (e *ItemPriceEntitlementsRemovedEvent) GetEventId() string       { return e.Id }
func (e *ItemPriceEntitlementsRemovedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SalesOrderUpdatedEvent represents a sales_order_updated webhook event
type SalesOrderUpdatedEvent struct {
	Id         string                    `json:"id"`
	OccurredAt time.Time                 `json:"occurred_at"`
	EventType  string                    `json:"event_type"`
	Content    *SalesOrderUpdatedContent `json:"content"`
}

func (e *SalesOrderUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *SalesOrderUpdatedEvent) GetEventId() string       { return e.Id }
func (e *SalesOrderUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// OmnichannelSubscriptionItemDunningStartedEvent represents a omnichannel_subscription_item_dunning_started webhook event
type OmnichannelSubscriptionItemDunningStartedEvent struct {
	Id         string                                            `json:"id"`
	OccurredAt time.Time                                         `json:"occurred_at"`
	EventType  string                                            `json:"event_type"`
	Content    *OmnichannelSubscriptionItemDunningStartedContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemDunningStartedEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// OmnichannelSubscriptionItemChangeScheduledEvent represents a omnichannel_subscription_item_change_scheduled webhook event
type OmnichannelSubscriptionItemChangeScheduledEvent struct {
	Id         string                                             `json:"id"`
	OccurredAt time.Time                                          `json:"occurred_at"`
	EventType  string                                             `json:"event_type"`
	Content    *OmnichannelSubscriptionItemChangeScheduledContent `json:"content"`
}

func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetEventType() string { return e.EventType }
func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetEventId() string   { return e.Id }
func (e *OmnichannelSubscriptionItemChangeScheduledEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

// PendingInvoiceUpdatedEvent represents a pending_invoice_updated webhook event
type PendingInvoiceUpdatedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *PendingInvoiceUpdatedContent `json:"content"`
}

func (e *PendingInvoiceUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *PendingInvoiceUpdatedEvent) GetEventId() string       { return e.Id }
func (e *PendingInvoiceUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// QuoteUpdatedEvent represents a quote_updated webhook event
type QuoteUpdatedEvent struct {
	Id         string               `json:"id"`
	OccurredAt time.Time            `json:"occurred_at"`
	EventType  string               `json:"event_type"`
	Content    *QuoteUpdatedContent `json:"content"`
}

func (e *QuoteUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *QuoteUpdatedEvent) GetEventId() string       { return e.Id }
func (e *QuoteUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// AttachedItemUpdatedEvent represents a attached_item_updated webhook event
type AttachedItemUpdatedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AttachedItemUpdatedContent `json:"content"`
}

func (e *AttachedItemUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *AttachedItemUpdatedEvent) GetEventId() string       { return e.Id }
func (e *AttachedItemUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// PaymentSourceUpdatedEvent represents a payment_source_updated webhook event
type PaymentSourceUpdatedEvent struct {
	Id         string                       `json:"id"`
	OccurredAt time.Time                    `json:"occurred_at"`
	EventType  string                       `json:"event_type"`
	Content    *PaymentSourceUpdatedContent `json:"content"`
}

func (e *PaymentSourceUpdatedEvent) GetEventType() string     { return e.EventType }
func (e *PaymentSourceUpdatedEvent) GetEventId() string       { return e.Id }
func (e *PaymentSourceUpdatedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// BusinessEntityDeletedEvent represents a business_entity_deleted webhook event
type BusinessEntityDeletedEvent struct {
	Id         string                        `json:"id"`
	OccurredAt time.Time                     `json:"occurred_at"`
	EventType  string                        `json:"event_type"`
	Content    *BusinessEntityDeletedContent `json:"content"`
}

func (e *BusinessEntityDeletedEvent) GetEventType() string     { return e.EventType }
func (e *BusinessEntityDeletedEvent) GetEventId() string       { return e.Id }
func (e *BusinessEntityDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// AuthorizationVoidedEvent represents a authorization_voided webhook event
type AuthorizationVoidedEvent struct {
	Id         string                      `json:"id"`
	OccurredAt time.Time                   `json:"occurred_at"`
	EventType  string                      `json:"event_type"`
	Content    *AuthorizationVoidedContent `json:"content"`
}

func (e *AuthorizationVoidedEvent) GetEventType() string     { return e.EventType }
func (e *AuthorizationVoidedEvent) GetEventId() string       { return e.Id }
func (e *AuthorizationVoidedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// SubscriptionRampDeletedEvent represents a subscription_ramp_deleted webhook event
type SubscriptionRampDeletedEvent struct {
	Id         string                          `json:"id"`
	OccurredAt time.Time                       `json:"occurred_at"`
	EventType  string                          `json:"event_type"`
	Content    *SubscriptionRampDeletedContent `json:"content"`
}

func (e *SubscriptionRampDeletedEvent) GetEventType() string     { return e.EventType }
func (e *SubscriptionRampDeletedEvent) GetEventId() string       { return e.Id }
func (e *SubscriptionRampDeletedEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// GenericWebhookEvent represents an unsupported webhook event
type GenericWebhookEvent struct {
	Id         string          `json:"id"`
	OccurredAt time.Time       `json:"occurred_at"`
	EventType  string          `json:"event_type"`
	RawContent json.RawMessage `json:"content"`
}

func (e *GenericWebhookEvent) GetEventType() string     { return e.EventType }
func (e *GenericWebhookEvent) GetEventId() string       { return e.Id }
func (e *GenericWebhookEvent) GetOccurredAt() time.Time { return e.OccurredAt }

// ParseWebhook deserializes webhook body and returns the appropriate typed event
func ParseWebhook(body string) (WebhookEventInterface, error) {
	event := Deserialize(body)

	// Parse content based on event type and return the specific typed event
	switch WebhookContentType(event.EventType.String()) {

	case WebhookContentTypeSubscriptionPauseScheduled:
		var content SubscriptionPauseScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionPauseScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerBusinessEntityChanged:
		var content CustomerBusinessEntityChangedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerBusinessEntityChangedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionAdvanceInvoiceScheduleAdded:
		var content SubscriptionAdvanceInvoiceScheduleAddedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionAdvanceInvoiceScheduleAddedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftExpired:
		var content GiftExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &GiftExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTaxWithheldDeleted:
		var content TaxWithheldDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TaxWithheldDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesDeleted:
		var content UnbilledChargesDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &UnbilledChargesDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponUpdated:
		var content CouponUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemReactivated:
		var content OmnichannelSubscriptionItemReactivatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemReactivatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemRenewed:
		var content OmnichannelSubscriptionItemRenewedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemRenewedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesCreated:
		var content UnbilledChargesCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &UnbilledChargesCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionResumed:
		var content SubscriptionResumedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionResumedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelOneTimeOrderItemCancelled:
		var content OmnichannelOneTimeOrderItemCancelledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelOneTimeOrderItemCancelledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCancelled:
		var content SubscriptionCancelledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionCancelledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemEntitlementsRemoved:
		var content ItemEntitlementsRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemEntitlementsRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeBusinessEntityCreated:
		var content BusinessEntityCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &BusinessEntityCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponSetUpdated:
		var content CouponSetUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponSetUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeDifferentialPriceUpdated:
		var content DifferentialPriceUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &DifferentialPriceUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemPaused:
		var content OmnichannelSubscriptionItemPausedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemPausedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeEntitlementOverridesRemoved:
		var content EntitlementOverridesRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &EntitlementOverridesRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionActivatedWithBackdating:
		var content SubscriptionActivatedWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionActivatedWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionTrialEndReminder:
		var content SubscriptionTrialEndReminderContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionTrialEndReminderEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionShippingAddressUpdated:
		var content SubscriptionShippingAddressUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionShippingAddressUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeVoucherCreateFailed:
		var content VoucherCreateFailedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &VoucherCreateFailedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftClaimed:
		var content GiftClaimedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &GiftClaimedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerDeleted:
		var content CustomerDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeRefundInitiated:
		var content RefundInitiatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &RefundInitiatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceGeneratedWithBackdating:
		var content InvoiceGeneratedWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &InvoiceGeneratedWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelTransactionCreated:
		var content OmnichannelTransactionCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelTransactionCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeAddUsagesReminder:
		var content AddUsagesReminderContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &AddUsagesReminderEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeVoucherCreated:
		var content VoucherCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &VoucherCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeRuleUpdated:
		var content RuleUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &RuleUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSchedulesCreated:
		var content PaymentSchedulesCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSchedulesCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureActivated:
		var content FeatureActivatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &FeatureActivatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceLocallyDeleted:
		var content PaymentSourceLocallyDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSourceLocallyDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceGenerated:
		var content InvoiceGeneratedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &InvoiceGeneratedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeVoucherExpired:
		var content VoucherExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &VoucherExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeAuthorizationSucceeded:
		var content AuthorizationSucceededContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &AuthorizationSucceededEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftScheduled:
		var content GiftScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &GiftScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionChangesScheduled:
		var content SubscriptionChangesScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionChangesScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionChangedWithBackdating:
		var content SubscriptionChangedWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionChangedWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemChanged:
		var content OmnichannelSubscriptionItemChangedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemChangedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftUnclaimed:
		var content GiftUnclaimedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &GiftUnclaimedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeVirtualBankAccountAdded:
		var content VirtualBankAccountAddedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &VirtualBankAccountAddedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentIntentCreated:
		var content PaymentIntentCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentIntentCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteCreatedWithBackdating:
		var content CreditNoteCreatedWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CreditNoteCreatedWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermTerminated:
		var content ContractTermTerminatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ContractTermTerminatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemFamilyUpdated:
		var content ItemFamilyUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemFamilyUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderCreated:
		var content OrderCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePriceVariantDeleted:
		var content PriceVariantDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PriceVariantDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionMovementFailed:
		var content SubscriptionMovementFailedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionMovementFailedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerMovedIn:
		var content CustomerMovedInContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerMovedInEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionAdvanceInvoiceScheduleUpdated:
		var content SubscriptionAdvanceInvoiceScheduleUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionAdvanceInvoiceScheduleUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemDeleted:
		var content ItemDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampDrafted:
		var content SubscriptionRampDraftedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRampDraftedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeDunningUpdated:
		var content DunningUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &DunningUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemEntitlementsUpdated:
		var content ItemEntitlementsUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemEntitlementsUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTokenConsumed:
		var content TokenConsumedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TokenConsumedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeHierarchyDeleted:
		var content HierarchyDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &HierarchyDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCancellationScheduled:
		var content SubscriptionCancellationScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionCancellationScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRenewed:
		var content SubscriptionRenewedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRenewedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureUpdated:
		var content FeatureUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &FeatureUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureDeleted:
		var content FeatureDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &FeatureDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemFamilyCreated:
		var content ItemFamilyCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemFamilyCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemScheduledChangeRemoved:
		var content OmnichannelSubscriptionItemScheduledChangeRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemScheduledChangeRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemResumed:
		var content OmnichannelSubscriptionItemResumedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemResumedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePurchaseCreated:
		var content PurchaseCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PurchaseCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeEntitlementOverridesUpdated:
		var content EntitlementOverridesUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &EntitlementOverridesUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemFamilyDeleted:
		var content ItemFamilyDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemFamilyDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionResumptionScheduled:
		var content SubscriptionResumptionScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionResumptionScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureReactivated:
		var content FeatureReactivatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &FeatureReactivatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCodesDeleted:
		var content CouponCodesDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponCodesDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardExpired:
		var content CardExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CardExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteUpdated:
		var content CreditNoteUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CreditNoteUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemDowngraded:
		var content OmnichannelSubscriptionItemDowngradedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemDowngradedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePriceVariantUpdated:
		var content PriceVariantUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PriceVariantUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePromotionalCreditsDeducted:
		var content PromotionalCreditsDeductedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PromotionalCreditsDeductedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampApplied:
		var content SubscriptionRampAppliedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRampAppliedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionPaused:
		var content SubscriptionPausedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionPausedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderReadyToProcess:
		var content OrderReadyToProcessContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderReadyToProcessEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureCreated:
		var content FeatureCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &FeatureCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTransactionDeleted:
		var content TransactionDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TransactionDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteCreated:
		var content CreditNoteCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CreditNoteCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemResubscribed:
		var content OmnichannelSubscriptionItemResubscribedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemResubscribedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeRecordPurchaseFailed:
		var content RecordPurchaseFailedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &RecordPurchaseFailedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemCreated:
		var content ItemCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTransactionUpdated:
		var content TransactionUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TransactionUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeMrrUpdated:
		var content MrrUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &MrrUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesInvoiced:
		var content UnbilledChargesInvoicedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &UnbilledChargesInvoicedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceUpdated:
		var content ItemPriceUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemPriceUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCodesUpdated:
		var content CouponCodesUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponCodesUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeVirtualBankAccountUpdated:
		var content VirtualBankAccountUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &VirtualBankAccountUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermCreated:
		var content ContractTermCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ContractTermCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionChanged:
		var content SubscriptionChangedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionChangedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentFailed:
		var content PaymentFailedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentFailedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCreditNoteDeleted:
		var content CreditNoteDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CreditNoteDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTaxWithheldRefunded:
		var content TaxWithheldRefundedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TaxWithheldRefundedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermCompleted:
		var content ContractTermCompletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ContractTermCompletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSchedulesUpdated:
		var content PaymentSchedulesUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSchedulesUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemExpired:
		var content OmnichannelSubscriptionItemExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardUpdated:
		var content CardUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CardUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerCreated:
		var content CustomerCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRenewalReminder:
		var content SubscriptionRenewalReminderContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRenewalReminderEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderDelivered:
		var content OrderDeliveredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderDeliveredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemCancellationScheduled:
		var content OmnichannelSubscriptionItemCancellationScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemCancellationScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemGracePeriodExpired:
		var content OmnichannelSubscriptionItemGracePeriodExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemGracePeriodExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCodesAdded:
		var content CouponCodesAddedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponCodesAddedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftCancelled:
		var content GiftCancelledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &GiftCancelledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderCancelled:
		var content OrderCancelledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderCancelledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponDeleted:
		var content CouponDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledChangesRemoved:
		var content SubscriptionScheduledChangesRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionScheduledChangesRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePendingInvoiceCreated:
		var content PendingInvoiceCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PendingInvoiceCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeEntitlementOverridesAutoRemoved:
		var content EntitlementOverridesAutoRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &EntitlementOverridesAutoRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemUpgraded:
		var content OmnichannelSubscriptionItemUpgradedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemUpgradedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionBusinessEntityChanged:
		var content SubscriptionBusinessEntityChangedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionBusinessEntityChangedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelOneTimeOrderCreated:
		var content OmnichannelOneTimeOrderCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelOneTimeOrderCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceDeleted:
		var content PaymentSourceDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSourceDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemCancelled:
		var content OmnichannelSubscriptionItemCancelledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemCancelledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeQuoteDeleted:
		var content QuoteDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &QuoteDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceUpdated:
		var content InvoiceUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &InvoiceUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionAdvanceInvoiceScheduleRemoved:
		var content SubscriptionAdvanceInvoiceScheduleRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionAdvanceInvoiceScheduleRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardDeleted:
		var content CardDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CardDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderReadyToShip:
		var content OrderReadyToShipContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderReadyToShipEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionMovedOut:
		var content SubscriptionMovedOutContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionMovedOutEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentScheduleSchemeCreated:
		var content PaymentScheduleSchemeCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentScheduleSchemeCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeBusinessEntityUpdated:
		var content BusinessEntityUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &BusinessEntityUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledResumptionRemoved:
		var content SubscriptionScheduledResumptionRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionScheduledResumptionRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentInitiated:
		var content PaymentInitiatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentInitiatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeFeatureArchived:
		var content FeatureArchivedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &FeatureArchivedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionReactivatedWithBackdating:
		var content SubscriptionReactivatedWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionReactivatedWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionImported:
		var content OmnichannelSubscriptionImportedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionImportedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTokenExpired:
		var content TokenExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TokenExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardAdded:
		var content CardAddedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CardAddedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponCreated:
		var content CouponCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeRuleDeleted:
		var content RuleDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &RuleDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceEntitlementsUpdated:
		var content ItemPriceEntitlementsUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemPriceEntitlementsUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceDeleted:
		var content ItemPriceDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemPriceDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeVirtualBankAccountDeleted:
		var content VirtualBankAccountDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &VirtualBankAccountDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentScheduleSchemeDeleted:
		var content PaymentScheduleSchemeDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentScheduleSchemeDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCreated:
		var content SubscriptionCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionEntitlementsCreated:
		var content SubscriptionEntitlementsCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionEntitlementsCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderReturned:
		var content OrderReturnedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderReturnedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionDeleted:
		var content SubscriptionDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceAdded:
		var content PaymentSourceAddedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSourceAddedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionMovedIn:
		var content SubscriptionMovedInContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionMovedInEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceCreated:
		var content ItemPriceCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemPriceCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledCancellationRemoved:
		var content SubscriptionScheduledCancellationRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionScheduledCancellationRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentRefunded:
		var content PaymentRefundedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentRefundedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeUsageFileIngested:
		var content UsageFileIngestedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &UsageFileIngestedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionMovedIn:
		var content OmnichannelSubscriptionMovedInContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionMovedInEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeDifferentialPriceCreated:
		var content DifferentialPriceCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &DifferentialPriceCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTransactionCreated:
		var content TransactionCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TransactionCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSucceeded:
		var content PaymentSucceededContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSucceededEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCanceledWithBackdating:
		var content SubscriptionCanceledWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionCanceledWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeUnbilledChargesVoided:
		var content UnbilledChargesVoidedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &UnbilledChargesVoidedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeQuoteCreated:
		var content QuoteCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &QuoteCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponSetDeleted:
		var content CouponSetDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponSetDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeAttachedItemCreated:
		var content AttachedItemCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &AttachedItemCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSalesOrderCreated:
		var content SalesOrderCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SalesOrderCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerChanged:
		var content CustomerChangedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerChangedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionStarted:
		var content SubscriptionStartedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionStartedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionActivated:
		var content SubscriptionActivatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionActivatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceExpiring:
		var content PaymentSourceExpiringContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSourceExpiringEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionReactivated:
		var content SubscriptionReactivatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionReactivatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderUpdated:
		var content OrderUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionScheduledPauseRemoved:
		var content SubscriptionScheduledPauseRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionScheduledPauseRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCancellationReminder:
		var content SubscriptionCancellationReminderContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionCancellationReminderEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionCreatedWithBackdating:
		var content SubscriptionCreatedWithBackdatingContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionCreatedWithBackdatingEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampCreated:
		var content SubscriptionRampCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRampCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderDeleted:
		var content OrderDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemPauseScheduled:
		var content OmnichannelSubscriptionItemPauseScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemPauseScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeGiftUpdated:
		var content GiftUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &GiftUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionTrialExtended:
		var content SubscriptionTrialExtendedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionTrialExtendedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemGracePeriodStarted:
		var content OmnichannelSubscriptionItemGracePeriodStartedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemGracePeriodStartedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCardExpiryReminder:
		var content CardExpiryReminderContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CardExpiryReminderEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTokenCreated:
		var content TokenCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TokenCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePromotionalCreditsAdded:
		var content PromotionalCreditsAddedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PromotionalCreditsAddedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampUpdated:
		var content SubscriptionRampUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRampUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerEntitlementsUpdated:
		var content CustomerEntitlementsUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerEntitlementsUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceExpired:
		var content PaymentSourceExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSourceExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCustomerMovedOut:
		var content CustomerMovedOutContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CustomerMovedOutEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionEntitlementsUpdated:
		var content SubscriptionEntitlementsUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionEntitlementsUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemDunningExpired:
		var content OmnichannelSubscriptionItemDunningExpiredContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemDunningExpiredEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeHierarchyCreated:
		var content HierarchyCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &HierarchyCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeAttachedItemDeleted:
		var content AttachedItemDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &AttachedItemDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemScheduledCancellationRemoved:
		var content OmnichannelSubscriptionItemScheduledCancellationRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemScheduledCancellationRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemUpdated:
		var content ItemUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeCouponSetCreated:
		var content CouponSetCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &CouponSetCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentIntentUpdated:
		var content PaymentIntentUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentIntentUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOrderResent:
		var content OrderResentContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OrderResentEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionCreated:
		var content OmnichannelSubscriptionCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeTaxWithheldRecorded:
		var content TaxWithheldRecordedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &TaxWithheldRecordedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePriceVariantCreated:
		var content PriceVariantCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PriceVariantCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeDifferentialPriceDeleted:
		var content DifferentialPriceDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &DifferentialPriceDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionItemsRenewed:
		var content SubscriptionItemsRenewedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionItemsRenewedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeRuleCreated:
		var content RuleCreatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &RuleCreatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermCancelled:
		var content ContractTermCancelledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ContractTermCancelledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeContractTermRenewed:
		var content ContractTermRenewedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ContractTermRenewedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeInvoiceDeleted:
		var content InvoiceDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &InvoiceDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeItemPriceEntitlementsRemoved:
		var content ItemPriceEntitlementsRemovedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &ItemPriceEntitlementsRemovedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSalesOrderUpdated:
		var content SalesOrderUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SalesOrderUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemDunningStarted:
		var content OmnichannelSubscriptionItemDunningStartedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemDunningStartedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeOmnichannelSubscriptionItemChangeScheduled:
		var content OmnichannelSubscriptionItemChangeScheduledContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &OmnichannelSubscriptionItemChangeScheduledEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePendingInvoiceUpdated:
		var content PendingInvoiceUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PendingInvoiceUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeQuoteUpdated:
		var content QuoteUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &QuoteUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeAttachedItemUpdated:
		var content AttachedItemUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &AttachedItemUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypePaymentSourceUpdated:
		var content PaymentSourceUpdatedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &PaymentSourceUpdatedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeBusinessEntityDeleted:
		var content BusinessEntityDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &BusinessEntityDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeAuthorizationVoided:
		var content AuthorizationVoidedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &AuthorizationVoidedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	case WebhookContentTypeSubscriptionRampDeleted:
		var content SubscriptionRampDeletedContent
		err := json.Unmarshal(event.Content, &content)
		if err != nil {
			return nil, err
		}
		return &SubscriptionRampDeletedEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			Content:    &content,
		}, nil

	default:
		// For unsupported event types, return generic event
		return &GenericWebhookEvent{
			Id:         event.Id,
			OccurredAt: event.OccurredAt,
			EventType:  event.EventType.String(),
			RawContent: event.Content,
		}, nil
	}
}
