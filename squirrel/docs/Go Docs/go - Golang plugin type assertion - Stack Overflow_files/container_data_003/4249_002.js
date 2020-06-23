    var edgeSupport_954856 = "f";
    var isIEBrowser_954856=false;
    var browserVersion_954856;

    function initiateNewRequest_954856(edgeSupport, html5Support) {
            var newUrl = "https://ads.everesttech.net/ads/mts/15699/4249?DFA_Click_Tracker=https%3A%2F%2Fadclick.g.doubleclick.net%2Fpcs%2Fclick%3Fxai%3DAKAOjsvgyaM55SOTRdtfFVWI5d7x0rqpTwtPSG6XdZSJrerNGgx0ZHOGYQZqFA9LwQirMUwZ5-PkxO14pFtJsKwFrSTYRPeEIXUgJ_OSBepSOskCChO7gSQb5v8y-YM6YEbri3AaxyrKqmmFW2cvZ_KtOGJa_HolU8A2vfCxSx4XWJJi7tz8PckM1ZzhWR07gVx8z7S1v3VlrhhgYUB8EPmDehyvIsQvpUxzRAd5aoZ3EpZKxWbN_129qvJw_ZJ9k7W_CjOHc2rmNKNDKgoJOU5OoEFg2EaXwkdAIQebz55ok-X-wl1v58cY%26sai%3DAMfl-YTy7fwX9-Kz8qrm3ZiEfPtaVmjkM9pt_w7tDlsPRGrucu2k6aZoEHQLsKo2NwO2t6tLPD3b0q3ICcx6AB4LT28ySnbXAxMXgs3z9LDcqYumuYvy1t2bJl0PxA9FiH72%26sig%3DCg0ArKJSzGI52nxRzwGoEAE%26urlfix%3D1%26adurl%3Dhttps%3A%2F%2Fadclick.g.doubleclick.net%2Fpcs%2Fclick%3Fxai%3DAKAOjstdWp642159suqFJg7r4zC83taN-I9vJalcECFWfSoYI3ELNxC9CWbf1Qvrfxc4-eB0Po5wjaXTAHAy-zLDNrOkpBJLaez1JO2KxKGO6HRszJVeyV562sTkwrpNtwiyCg9OP7CGaF-GZlOiZ2VRSIoz2Uj6IKFVgR60xtbF-H-Q6kZOmWnGIPom1FWHwj_yhDbQNpqNPVX41SQ4_eaFi7Vmu1ZDhLHrJEUzrmq0lUNKhuSvnSD-K_jNGZv91sOgxAHxllqorsWnL4JTgXfMYeZV7oLyh_OpWNBmdOobp3UpPCvaDimUGy9hBgN8juCZRkHrPshivq2nBu2Zc65LsZuV4T-NJuze0GSJ-pXb2HSC3BFr5ml4X1RcSb6ribJgJrpYNTRDDCDwvoLdioe2WsGEQeyaqTeorSgODtmC9O8zHDGjECNO6BM9PvP3BRpd9CmtNbSn2eH14H0t7vE7eXLOkw87aZbjSKFv1bDm852xnHdVROEjMxk9oohUg6RjM1We2-bIPcxUm15bvsDel8wgiR97uvE5m9F1DIGfAGtpgOUqjBCYH8ni33nAOfUp8ltlLCqz4B02k5RoDCYAeJwlhNhHESuK8zZeYpdNllbHib1CYvT2pAmPAeG33e1PnJLFywqYnoY6VF9kb4ohDmlu9M0kURKQUJzzkSFltjHSvSNZ8lHAqI0o47ZOeC5vRXo5rUMNAqmkZeiU-5FyoVrQd0X9d9EVeH8XmltlHhxpw98ioDwcEXWNIqET-oWJjVyosefLf0AFPX8erUWAwCVzhwAjIPBSYKOLKt5xzLxP_l4L6dWpJphQ7ZSFKdSoBqHnzhKoE3AdnhLV_B6L0Ts7EXAag_eaBur8LW462sv6X9I7rkc6wmo5quvZ_kCnOvCKVcxfS4ftCqFyKrB7cIfhWuE3F-7Nfqnmvm-47d6bXqY_fcbn-wNp-GC0O-E_lo4C-YLptHuI1islsbvXrcpVQAhszxPBe5rZkhW6GrdQ9Y73MBUITtd4B1aqmMkd4JKcvUvz8XKnOK3TRRYwNwudfzebsIZ5_LTcooUbpXMv37uE%26sai%3DAMfl-YQTdKGCa2lkoXIA4MXWJYwkeHU6xWvnjbaRhWKc7AIfGsTnUHZ0LULo5RfYqlrHWGyUbtnRvVuN3xXZoIn_o5Ihpw42KrpBb0rRlc1HGNcVcNjPXjjSboFOWPLR6IV1GddcoB22mwEHVf-mqmGY4VEg72O_DlUt2_UPnCzy7lvIajBEI9W6qukGTBmKHrPbeqy5BsP3dhMJuZfNQx1O_6p4yNiktZpq2SICqtJHQrK_jZcR4A1-%26sig%3DCg0ArKJSzDtOunZINtxzEAE%26urlfix%3D1%26adurl%3D&DFA_BuyId=23926483&DFA_PlacementId=270363406&DFA_AdId=465079040&DFA_CreativeId=110158953&DFA_SiteId=3654125&TC_1=2000053&TC_2=23926483&TC_3=270363406&TC_4=110158953&TC_5=dcmadvertiserid|8391437$dcmcampaignid|23926483$dcmadid|465079040$dcmrenderingid|110264067$dcmsiteid|3654125$dcmplacementid|270363406$customer|Microsoft$dv360auctionid|ct=AT&st=&city=0&dma=0&zp=&bw=0&DCM_PlacementID=270363406" + "&edge=" + edgeSupport + "&html5="+ html5Support +"&nr=" + Math.random();
            if(document.readyState === "complete")
            {
                var sc = document.createElement("script");
                sc.setAttribute("type","text/javascript");
                sc.setAttribute("src",newUrl);
                if (document.currentScript) {
                    var pn = document.currentScript.parentNode;
                    var sbn = document.currentScript.nextSibling;
                    if (sbn) {
                        pn.insertBefore(sc,sbn);
                    } else {
                        pn.appendChild(sc);
                    }
                } else {
                    document.body.appendChild(sc);
                }
            } else {
                document.write('<' + 'script type="text/javascript" src="' + newUrl +'"></' + 'script>');
            }
        }

     function getInternetExplorerVersion_954856() {
         // Returns the version of Internet Explorer or a -1
         // (indicating the use of another browser).

             var rv = -1; // Return value assumes failure.
             if (navigator.appName == 'Microsoft Internet Explorer') {
                 isIEBrowser_954856=true;
                 var ua = navigator.userAgent;
                 var re  = new RegExp("MSIE ([0-9]{1,}[\.0-9]{0,})");

                 if (re.exec(ua) != null)
                     rv = parseFloat( RegExp.$1 );
             }

             return rv;
         }

      //returns true if ie version is less than 9, say ie6, ie7, ie8
         // -1 for non IE browsers.
         function isIEBrowserWithVersionLessThan9_954856 () {

             browserVersion_954856 = getInternetExplorerVersion_954856();  //-1 for non IE browsers
             if((browserVersion_954856 != -1) && (browserVersion_954856 < 9)) {
                 return true;

             }
             return false;
         }

    //code to detect Edge Features, courtesy  (http://dl.dropboxusercontent.com/u/13483458/test-edge.html)
    var testEle_954856=document.createElement("div_954856");
    function isSupported_954856(a){

        var d=testEle_954856.style,e;
        for(i=0;i<a.length;i++)
            if(e=a[i],d[e]!==void 0)
                return!0;
        return!1
    }

    function supportsRGBA_954856(){

        testEle_954856.cssText="background-color:rgba(150,255,150,.5)";
        if((""+testEle_954856.style.backgroundColor).indexOf("rgba")==0)
            return!0;
        return!1
    }

    var hasTransform_954856=isSupported_954856([
        "transformProperty",
        "WebkitTransform",
        "MozTransform",
        "OTransform",
        "msTransform"
    ]),

    hasSVG_954856=!!document.createElementNS&&!!document.createElementNS("http://www.w3.org/2000/svg","svg").createSVGRect,
    hasRGBA_954856=supportsRGBA_954856(),
    hasJSON_954856=window.JSON&&window.JSON.parse&&window.JSON.stringify,
    readyToPlay=!1;

    function isIEBrowserVersion9_954856() {
        return (isIEBrowser_954856 && (browserVersion_954856 == 9)) ? true : false;
    }

    function isEdgeSupported_954856() {
        if(isIEBrowserVersion9_954856()) {
            return "y";           //hardcoding IE9 edge support.
        }
        if(hasTransform_954856) {
            if(requiresSVG_954856&&!hasSVG_954856)
                return "f";
            return "y";
        }
        return "f";
    }

    function isCanvasSupported_954856(){
      var elem = document.createElement('canvas');
      return !!(elem.getContext && elem.getContext('2d'));
    }

    function isHTML5FeaturesSupported_954856() {
         return (isCanvasSupported_954856()) ? "y" : "f";
    }

    var requiresSVG_954856=false;
    //edge detection code end

    //Edge is not supported in IE 6,7,8. Hence hardcoding edge as not supported for the same.
   // edgeSupport_954856 = (isIEBrowserWithVersionLessThan9_954856()) ? "f" : isHTMLFeaturesSupported_954856(featureArray_954856);
    edgeSupport_954856 = (isIEBrowserWithVersionLessThan9_954856()) ? "f" : isEdgeSupported_954856();
    html5Support_954856 = isHTML5FeaturesSupported_954856();

    initiateNewRequest_954856(edgeSupport_954856, html5Support_954856);
