[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/jadiunr/check-netstat)
![Go Test](https://github.com/jadiunr/check-netstat/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/jadiunr/check-netstat/workflows/goreleaser/badge.svg)

# Sensu netstat check

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
  - [Check Output](#check-output)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The Sensu netstat check is a [Sensu Check][6] that reports network protocol statistics.

## Usage examples

```
Simple cross-platform network statistics checks

Usage:
  check-netstat [flags]
  check-netstat [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help   help for check-netstat

Use "check-netstat [command] --help" for more information about a command.
```

### Check Output

<details>
<summary>spoiler</summary>

```
# HELP icmp_in_addr_mask_reps Statistic icmp InAddrMaskReps
# TYPE icmp_in_addr_mask_reps untyped
icmp_in_addr_mask_reps 0
# HELP icmp_in_addr_masks Statistic icmp InAddrMasks
# TYPE icmp_in_addr_masks untyped
icmp_in_addr_masks 0
# HELP icmp_in_csum_errors Statistic icmp InCsumErrors
# TYPE icmp_in_csum_errors untyped
icmp_in_csum_errors 0
# HELP icmp_in_dest_unreachs Statistic icmp InDestUnreachs
# TYPE icmp_in_dest_unreachs untyped
icmp_in_dest_unreachs 202
# HELP icmp_in_echo_reps Statistic icmp InEchoReps
# TYPE icmp_in_echo_reps untyped
icmp_in_echo_reps 0
# HELP icmp_in_echos Statistic icmp InEchos
# TYPE icmp_in_echos untyped
icmp_in_echos 0
# HELP icmp_in_errors Statistic icmp InErrors
# TYPE icmp_in_errors untyped
icmp_in_errors 0
# HELP icmp_in_msgs Statistic icmp InMsgs
# TYPE icmp_in_msgs untyped
icmp_in_msgs 202
# HELP icmp_in_parm_probs Statistic icmp InParmProbs
# TYPE icmp_in_parm_probs untyped
icmp_in_parm_probs 0
# HELP icmp_in_redirects Statistic icmp InRedirects
# TYPE icmp_in_redirects untyped
icmp_in_redirects 0
# HELP icmp_in_src_quenchs Statistic icmp InSrcQuenchs
# TYPE icmp_in_src_quenchs untyped
icmp_in_src_quenchs 0
# HELP icmp_in_time_excds Statistic icmp InTimeExcds
# TYPE icmp_in_time_excds untyped
icmp_in_time_excds 0
# HELP icmp_in_timestamp_reps Statistic icmp InTimestampReps
# TYPE icmp_in_timestamp_reps untyped
icmp_in_timestamp_reps 0
# HELP icmp_in_timestamps Statistic icmp InTimestamps
# TYPE icmp_in_timestamps untyped
icmp_in_timestamps 0
# HELP icmp_out_addr_mask_reps Statistic icmp OutAddrMaskReps
# TYPE icmp_out_addr_mask_reps untyped
icmp_out_addr_mask_reps 0
# HELP icmp_out_addr_masks Statistic icmp OutAddrMasks
# TYPE icmp_out_addr_masks untyped
icmp_out_addr_masks 0
# HELP icmp_out_dest_unreachs Statistic icmp OutDestUnreachs
# TYPE icmp_out_dest_unreachs untyped
icmp_out_dest_unreachs 202
# HELP icmp_out_echo_reps Statistic icmp OutEchoReps
# TYPE icmp_out_echo_reps untyped
icmp_out_echo_reps 0
# HELP icmp_out_echos Statistic icmp OutEchos
# TYPE icmp_out_echos untyped
icmp_out_echos 0
# HELP icmp_out_errors Statistic icmp OutErrors
# TYPE icmp_out_errors untyped
icmp_out_errors 0
# HELP icmp_out_msgs Statistic icmp OutMsgs
# TYPE icmp_out_msgs untyped
icmp_out_msgs 202
# HELP icmp_out_parm_probs Statistic icmp OutParmProbs
# TYPE icmp_out_parm_probs untyped
icmp_out_parm_probs 0
# HELP icmp_out_redirects Statistic icmp OutRedirects
# TYPE icmp_out_redirects untyped
icmp_out_redirects 0
# HELP icmp_out_src_quenchs Statistic icmp OutSrcQuenchs
# TYPE icmp_out_src_quenchs untyped
icmp_out_src_quenchs 0
# HELP icmp_out_time_excds Statistic icmp OutTimeExcds
# TYPE icmp_out_time_excds untyped
icmp_out_time_excds 0
# HELP icmp_out_timestamp_reps Statistic icmp OutTimestampReps
# TYPE icmp_out_timestamp_reps untyped
icmp_out_timestamp_reps 0
# HELP icmp_out_timestamps Statistic icmp OutTimestamps
# TYPE icmp_out_timestamps untyped
icmp_out_timestamps 0
# HELP icmpmsg_in_type_3 Statistic icmpmsg InType3
# TYPE icmpmsg_in_type_3 untyped
icmpmsg_in_type_3 202
# HELP icmpmsg_out_type_3 Statistic icmpmsg OutType3
# TYPE icmpmsg_out_type_3 untyped
icmpmsg_out_type_3 202
# HELP ip_default_ttl Statistic ip DefaultTTL
# TYPE ip_default_ttl untyped
ip_default_ttl 64
# HELP ip_forw_datagrams Statistic ip ForwDatagrams
# TYPE ip_forw_datagrams untyped
ip_forw_datagrams 0
# HELP ip_forwarding Statistic ip Forwarding
# TYPE ip_forwarding untyped
ip_forwarding 1
# HELP ip_frag_creates Statistic ip FragCreates
# TYPE ip_frag_creates untyped
ip_frag_creates 4
# HELP ip_frag_fails Statistic ip FragFails
# TYPE ip_frag_fails untyped
ip_frag_fails 0
# HELP ip_frag_o_ks Statistic ip FragOKs
# TYPE ip_frag_o_ks untyped
ip_frag_o_ks 2
# HELP ip_in_addr_errors Statistic ip InAddrErrors
# TYPE ip_in_addr_errors untyped
ip_in_addr_errors 0
# HELP ip_in_delivers Statistic ip InDelivers
# TYPE ip_in_delivers untyped
ip_in_delivers 44240
# HELP ip_in_discards Statistic ip InDiscards
# TYPE ip_in_discards untyped
ip_in_discards 0
# HELP ip_in_hdr_errors Statistic ip InHdrErrors
# TYPE ip_in_hdr_errors untyped
ip_in_hdr_errors 0
# HELP ip_in_receives Statistic ip InReceives
# TYPE ip_in_receives untyped
ip_in_receives 44243
# HELP ip_in_unknown_protos Statistic ip InUnknownProtos
# TYPE ip_in_unknown_protos untyped
ip_in_unknown_protos 0
# HELP ip_out_discards Statistic ip OutDiscards
# TYPE ip_out_discards untyped
ip_out_discards 99
# HELP ip_out_no_routes Statistic ip OutNoRoutes
# TYPE ip_out_no_routes untyped
ip_out_no_routes 1
# HELP ip_out_requests Statistic ip OutRequests
# TYPE ip_out_requests untyped
ip_out_requests 32106
# HELP ip_reasm_fails Statistic ip ReasmFails
# TYPE ip_reasm_fails untyped
ip_reasm_fails 0
# HELP ip_reasm_o_ks Statistic ip ReasmOKs
# TYPE ip_reasm_o_ks untyped
ip_reasm_o_ks 0
# HELP ip_reasm_reqds Statistic ip ReasmReqds
# TYPE ip_reasm_reqds untyped
ip_reasm_reqds 0
# HELP ip_reasm_timeout Statistic ip ReasmTimeout
# TYPE ip_reasm_timeout untyped
ip_reasm_timeout 0
# HELP tcp_active_opens Statistic tcp ActiveOpens
# TYPE tcp_active_opens untyped
tcp_active_opens 275
# HELP tcp_attempt_fails Statistic tcp AttemptFails
# TYPE tcp_attempt_fails untyped
tcp_attempt_fails 0
# HELP tcp_curr_estab Statistic tcp CurrEstab
# TYPE tcp_curr_estab untyped
tcp_curr_estab 35
# HELP tcp_estab_resets Statistic tcp EstabResets
# TYPE tcp_estab_resets untyped
tcp_estab_resets 35
# HELP tcp_in_csum_errors Statistic tcp InCsumErrors
# TYPE tcp_in_csum_errors untyped
tcp_in_csum_errors 0
# HELP tcp_in_errs Statistic tcp InErrs
# TYPE tcp_in_errs untyped
tcp_in_errs 1
# HELP tcp_in_segs Statistic tcp InSegs
# TYPE tcp_in_segs untyped
tcp_in_segs 28344
# HELP tcp_max_conn Statistic tcp MaxConn
# TYPE tcp_max_conn untyped
tcp_max_conn -1
# HELP tcp_out_rsts Statistic tcp OutRsts
# TYPE tcp_out_rsts untyped
tcp_out_rsts 171
# HELP tcp_out_segs Statistic tcp OutSegs
# TYPE tcp_out_segs untyped
tcp_out_segs 28984
# HELP tcp_passive_opens Statistic tcp PassiveOpens
# TYPE tcp_passive_opens untyped
tcp_passive_opens 14
# HELP tcp_retrans_segs Statistic tcp RetransSegs
# TYPE tcp_retrans_segs untyped
tcp_retrans_segs 36
# HELP tcp_rto_algorithm Statistic tcp RtoAlgorithm
# TYPE tcp_rto_algorithm untyped
tcp_rto_algorithm 1
# HELP tcp_rto_max Statistic tcp RtoMax
# TYPE tcp_rto_max untyped
tcp_rto_max 120000
# HELP tcp_rto_min Statistic tcp RtoMin
# TYPE tcp_rto_min untyped
tcp_rto_min 200
# HELP udp_ignored_multi Statistic udp IgnoredMulti
# TYPE udp_ignored_multi untyped
udp_ignored_multi 2902
# HELP udp_in_csum_errors Statistic udp InCsumErrors
# TYPE udp_in_csum_errors untyped
udp_in_csum_errors 0
# HELP udp_in_datagrams Statistic udp InDatagrams
# TYPE udp_in_datagrams untyped
udp_in_datagrams 19206
# HELP udp_in_errors Statistic udp InErrors
# TYPE udp_in_errors untyped
udp_in_errors 0
# HELP udp_mem_errors Statistic udp MemErrors
# TYPE udp_mem_errors untyped
udp_mem_errors 0
# HELP udp_no_ports Statistic udp NoPorts
# TYPE udp_no_ports untyped
udp_no_ports 202
# HELP udp_out_datagrams Statistic udp OutDatagrams
# TYPE udp_out_datagrams untyped
udp_out_datagrams 6313
# HELP udp_rcvbuf_errors Statistic udp RcvbufErrors
# TYPE udp_rcvbuf_errors untyped
udp_rcvbuf_errors 0
# HELP udp_sndbuf_errors Statistic udp SndbufErrors
# TYPE udp_sndbuf_errors untyped
udp_sndbuf_errors 1
# HELP udplite_ignored_multi Statistic udplite IgnoredMulti
# TYPE udplite_ignored_multi untyped
udplite_ignored_multi 0
# HELP udplite_in_csum_errors Statistic udplite InCsumErrors
# TYPE udplite_in_csum_errors untyped
udplite_in_csum_errors 0
# HELP udplite_in_datagrams Statistic udplite InDatagrams
# TYPE udplite_in_datagrams untyped
udplite_in_datagrams 0
# HELP udplite_in_errors Statistic udplite InErrors
# TYPE udplite_in_errors untyped
udplite_in_errors 0
# HELP udplite_mem_errors Statistic udplite MemErrors
# TYPE udplite_mem_errors untyped
udplite_mem_errors 0
# HELP udplite_no_ports Statistic udplite NoPorts
# TYPE udplite_no_ports untyped
udplite_no_ports 0
# HELP udplite_out_datagrams Statistic udplite OutDatagrams
# TYPE udplite_out_datagrams untyped
udplite_out_datagrams 0
# HELP udplite_rcvbuf_errors Statistic udplite RcvbufErrors
# TYPE udplite_rcvbuf_errors untyped
udplite_rcvbuf_errors 0
# HELP udplite_sndbuf_errors Statistic udplite SndbufErrors
# TYPE udplite_sndbuf_errors untyped
udplite_sndbuf_errors 0
```

</details>

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add jadiunr/check-netstat
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/jadiunr/check-netstat].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-netstat
  namespace: default
spec:
  command: check-netstat
  subscriptions:
  - system
  runtime_assets:
  - jadiunr/check-netstat
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the check-netstat repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/check-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/check-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu-community/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
