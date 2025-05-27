package questionbank

import "fmt"

// BinaryFindTarget
// 给定一个长度为n的有序数组 nums ，其中可能包含重复元素。请返回数组中最左一个元素 target 的索引。若数组中不包含该元素，则返回-1.
func BinaryFindTarget(nums []int, target int) int {
	i, j := 0, len(nums)-1
	mid := (j-i)/2 + i
	for i <= j {
		fmt.Println("mid: ", mid)
		if nums[mid] == target {
			left := mid
			for left >= 0 && nums[left] == target {
				left--
			}
			return left + 1
		}
		// <---
		if nums[mid] > target {
			j = mid - 1
			mid = (j-i)/2 + i
		}
		// --->
		if nums[mid] < target {
			i = mid + 1
			mid = (j-i)/2 + i
		}
	}
	return -1
}

// curl -v --location 'https://space.shopee.io/api/bromo/v1/service/get?env=live&service_name=edge-swc-live-id&detail=true' \
// --header 'accept: application/json' \
// --header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Ijc1YTIyMzA3NzdjNzY1YzYxOGUxOTFlZjY4NDU0MDhmN2JkYjY4MTQiLCJ0eXAiOiJKV1QifQ.eyJlbWFpbCI6ImRpbmd5dWFuLnd1QHNob3BlZS5jb20iLCJleHAiOjE3NDI0Mzg4NjgsImlzcyI6Imh0dHBzOi8vc3BhY2Uuc2hvcGVlLmlvL3YxL2NlcnRzIiwic3ViIjoiNDYwOTkzIn0.cG1ovimImIY6z4mIfD64ZQWP5j5wzBO7PqSAyE4_mbgfdhYdCNaCudegdn7qMr-Gm-TQcBXHqO-n-3kVUvzWXffOewnobjTwUbVV9DTy4J6xe9HA5-n8ZOy5IyRBzlkNQts5wK1K3VBnX7ODw5Osvz7iwlEVLsLY5GO6yGQOGyLgHrOWNxnUojXBD2XcOByR0hDeHy9ETQ-A0aOaCU84H1-Bn63_nAT_kXokiLtRfTcT990XRK8EaDZPM3rtnlHlHKIRFh6UGOTFpkeOhMmHs6bkcQJ34S4d4Q96qNf802yhRvEANdKJZ7XGGyMPDMAjWBBpIeMNZY_nM-GMMYztg2iXgST6tQKXvWB43rgNiRCzoP_uXrUUGzt49p_MVnaUuFKoPgzx89X4c8fDuCEBfvZ85AXsZECVJ589mVny19RP2vSURSjanEbXtj086EhpcQNctB1DfIM3aOp-XVkmXMBSPhvzszwqgTDJZfXgeOEGsf8ij1ALkBGFv0hwoMpeCq1nziZqUl3DYLBu3_VEdBzGj_BCuuqCV9Ia3vWU3h_3z5wnR2wmnjfoJpv9a8ENesZsUAg_ulIeGzl7XHl2h-2d7-O4u4j674aDfHMLdgWKvllmSUmuT4j0pi47YpfleS53kjeEtkKb7Y5mPee1XRRPBUF7v-8tDngx0GPbK5Q'

// curl -o /dev/null -s -w "DNS解析: %{time_namelookup}s\n连接建立: %{time_connect}s\nTLS握手: %{time_appconnect}s\n请求发送: %{time_pretransfer}s\n服务器响应: %{time_starttransfer}s\n总耗时: %{time_total}s\n"  \
// --location 'http://10.63.80.22:8081/api/deployment/list_brief?env=live&service_name=edge-swc-live-id&detail=true' \
// --header 'accept: application/json' \
// --header 'X-Forwarded-Email: bromo_meta_bot@shopee.com'

// curl --location 'http://10.63.80.22:8081/api/deployment/list_brief?env=live&service_name=edge-swc-live-id&detail=true' \
// --header 'accept: application/json' \
// --header 'X-Forwarded-Email: bromo_meta_bot@shopee.com'
