    var edgeSupport_985789 = "f";
    var isIEBrowser_985789=false;
    var browserVersion_985789;

    function initiateNewRequest_985789(edgeSupport, html5Support) {
            var newUrl = "https://ads.everesttech.net/ads/mts/15700/4249?DFA_Click_Tracker=https%3A%2F%2Fadclick.g.doubleclick.net%2Fpcs%2Fclick%3Fxai%3DAKAOjsvKJd_OJRAtMA10l8T4MqWtF79w_CKR7_8hPhMTkDMyvOc44A482sgUvWhrrM2aRq0Lss29Cn7BI6deN7E2tysjZaEhYtr4p5QtFRi9sdWH3pjpaitC8Snx8AhrHJC_9xNMV8tOQzKegaj49SMn1nX74sn8gnbIN0TnHZcdkE8RFQUAqS-IBxejN9fw0UZUgwLW_cz4tNHqEKEDWMCVNl3dJIf3jAm1Thr2aczdRZ5c1IWKI1JyOAUjUZaPOsACSGs6-BiLBLxUyhhqA7xtCEd0ep8gugMrCmPEstrmFmLtxDkjGAC5%26sai%3DAMfl-YRXqWhpjMEXRydKCuwgr4uOKjC2WvH34vN0tdNPIdijtZoinYImAJiOezpLE7-TuPNvaoDIXSw7nA2Y_uJ7vLUtjJ01kRDkYYe9DHnaDR7WQhfzGmhuHfQ2tr8gl8XU%26sig%3DCg0ArKJSzGzGE2USZ5PtEAE%26urlfix%3D1%26adurl%3Dhttps%3A%2F%2Fadclick.g.doubleclick.net%2Fpcs%2Fclick%3Fxai%3DAKAOjsucYQ-LoAL20ujIrqKHZ9wdZacgVkE-T_XmhVEWRp4dHl65LgjfVRBPM44agVOi55_Dv-85sqLo-hkeDcqPlqYlaDTNGfaz0BAkgoU7-tfUniY21_IZtV4reBdFooJIHD9b0ucakaU500-2gZotiUoc9iU42jca41g9j1oFuRrtG7A_hYinfXd2n5ZuSTwpfNArD2CyQjp0uBqZLBKXhBR5PdD_DrIye3j7_MsaawNnXCH5ZALXLq6KLXQsy2uYaef1VRVvO6oE-Wms5kWTk7OGzh0VKlVY4EHQGx9Kc5cdZpLpgoysU79Yr7Z-FaUCONaSFhNS_Og7acIskA0_teozm32NMS15kJ9Vu-Z4fv0YdWpsyE3foHhr-YnYA-uF4uIvMgJz5kfZfDO_U97fOoFVg2V5uZkGuv_mly4msnjrMrwa2ifleS8T-fx0Un1-8cMWD2RyrPIvTkYfIZC5fgmiXHexZNtPNYwGUDUhHMQI6HHeZxt3T09gFNpkJ8yKpQyJPwQqKUPoZ872ENWQC_iptCurEUo3PcAGDoiVVKvdamuNThrZ_mLQTvJ0WCp7AidH2WFUvyAmmVomUI-zTCL3_GsE3xAC8WXPvUUdAq1ccWXxVUiQHJI_X_GrZXeDdw7LBHf-G1XZC1owb-zG8F32kqd-AYcdTtr8xhJgMxHQhTwBpehyEsctdIkFH3CnpT6LshRFcsO7Dqwphq1aTDwYOjwWeiTV-bALGJousBXCDf8CaAqUeRBk5Fnpp9oe-7AFHrhp0gATUbPdqhkO2QlpuCiN879tcF12GUv3fgFYdXcBkzLvnmXnHl9k72F9lJYMuRluyjknjDMkJIHZ6cGeQEGosczsecK6cLAfay5FZ_A0K1L2WGqWPvT__oYpBQ4je8B8E65GrT2rKwUucvEEFStS6EraRo9-kzUoy66VfbO8Pfx6wVbVO2XsJmxo1RFMhX5-jePH5_atDa1dc7UcCTE2uigE3FPD9tXa5L-_AJ6EZSCy6Xt229xynHogO4iEg3r_iTJKz54Z4SwyKlpkTCm_Ub7cFGgXRGaEzH40OAM%26sai%3DAMfl-YQm4xB-oQ_tziarUaL8lLi1Kz4furPU5ZCL18DPKKfk8fVdywGFZoB4s--6paOLnOOTHFqhDd6GxwSGovGZTf72-nb0nGjTqz-K01WvFenan5NxH3tpP9F7VOPPkl-Q72LKDWdKBjcoDgKpdv9s88N_nhnrGO-xk9rwSklT8yCxoxl7YbPYLYZGY3AYysNSbcs0A0en99p7YPbceOjtQCCTJYw0kXgLiIXeSgBg7jyxhmbRWquk%26sig%3DCg0ArKJSzH1eBaiBrn0jEAE%26urlfix%3D1%26adurl%3D&DFA_BuyId=23926483&DFA_PlacementId=270097097&DFA_AdId=465069236&DFA_CreativeId=110158371&DFA_SiteId=3654125&TC_1=2000053&TC_2=23926483&TC_3=270097097&TC_4=110158371&TC_5=dcmadvertiserid|8391437$dcmcampaignid|23926483$dcmadid|465069236$dcmrenderingid|110262885$dcmsiteid|3654125$dcmplacementid|270097097$customer|Microsoft$dv360auctionid|ct=AT&st=&city=0&dma=0&zp=&bw=0&DCM_PlacementID=270097097" + "&edge=" + edgeSupport + "&html5="+ html5Support +"&nr=" + Math.random();
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

     function getInternetExplorerVersion_985789() {
         // Returns the version of Internet Explorer or a -1
         // (indicating the use of another browser).

             var rv = -1; // Return value assumes failure.
             if (navigator.appName == 'Microsoft Internet Explorer') {
                 isIEBrowser_985789=true;
                 var ua = navigator.userAgent;
                 var re  = new RegExp("MSIE ([0-9]{1,}[\.0-9]{0,})");

                 if (re.exec(ua) != null)
                     rv = parseFloat( RegExp.$1 );
             }

             return rv;
         }

      //returns true if ie version is less than 9, say ie6, ie7, ie8
         // -1 for non IE browsers.
         function isIEBrowserWithVersionLessThan9_985789 () {

             browserVersion_985789 = getInternetExplorerVersion_985789();  //-1 for non IE browsers
             if((browserVersion_985789 != -1) && (browserVersion_985789 < 9)) {
                 return true;

             }
             return false;
         }

    //code to detect Edge Features, courtesy  (http://dl.dropboxusercontent.com/u/13483458/test-edge.html)
    var testEle_985789=document.createElement("div_985789");
    function isSupported_985789(a){

        var d=testEle_985789.style,e;
        for(i=0;i<a.length;i++)
            if(e=a[i],d[e]!==void 0)
                return!0;
        return!1
    }

    function supportsRGBA_985789(){

        testEle_985789.cssText="background-color:rgba(150,255,150,.5)";
        if((""+testEle_985789.style.backgroundColor).indexOf("rgba")==0)
            return!0;
        return!1
    }

    var hasTransform_985789=isSupported_985789([
        "transformProperty",
        "WebkitTransform",
        "MozTransform",
        "OTransform",
        "msTransform"
    ]),

    hasSVG_985789=!!document.createElementNS&&!!document.createElementNS("http://www.w3.org/2000/svg","svg").createSVGRect,
    hasRGBA_985789=supportsRGBA_985789(),
    hasJSON_985789=window.JSON&&window.JSON.parse&&window.JSON.stringify,
    readyToPlay=!1;

    function isIEBrowserVersion9_985789() {
        return (isIEBrowser_985789 && (browserVersion_985789 == 9)) ? true : false;
    }

    function isEdgeSupported_985789() {
        if(isIEBrowserVersion9_985789()) {
            return "y";           //hardcoding IE9 edge support.
        }
        if(hasTransform_985789) {
            if(requiresSVG_985789&&!hasSVG_985789)
                return "f";
            return "y";
        }
        return "f";
    }

    function isCanvasSupported_985789(){
      var elem = document.createElement('canvas');
      return !!(elem.getContext && elem.getContext('2d'));
    }

    function isHTML5FeaturesSupported_985789() {
         return (isCanvasSupported_985789()) ? "y" : "f";
    }

    var requiresSVG_985789=false;
    //edge detection code end

    //Edge is not supported in IE 6,7,8. Hence hardcoding edge as not supported for the same.
   // edgeSupport_985789 = (isIEBrowserWithVersionLessThan9_985789()) ? "f" : isHTMLFeaturesSupported_985789(featureArray_985789);
    edgeSupport_985789 = (isIEBrowserWithVersionLessThan9_985789()) ? "f" : isEdgeSupported_985789();
    html5Support_985789 = isHTML5FeaturesSupported_985789();

    initiateNewRequest_985789(edgeSupport_985789, html5Support_985789);
