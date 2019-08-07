package samplestructs

import "encoding/xml"

// Server xml response:

//<servers version="1">
//     <server>
//         <serverName>Shanghai_VPN</serverName>
//         <serverIP>127.0.0.1</serverIP>
//     </server>
//     <server>
//         <serverName>Beijing_VPN</serverName>
//         <serverIP>127.0.0.2</serverIP>
//     </server>
// </servers>

// Structs for the above response:
type servers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

// Sample authentication response:

// <rsp stat="ok" version="1.0">
//     <api_key>6cae215092144319c0cd4b7667575240</api_key>
//     <version>4</version>
// </rsp>

// Struct for the above response:
type xmlResp struct {
	XMLName    xml.Name `xml:"rsp"`
	Version    string   `xml:"version,attr"`
	APIKey     string   `xml:"api_key"`
	VersionTag string   `xml:"version"`
}

// Sample visitor response:
// <rsp stat="ok" version="1.0">
//     <result>
//         <total_results>443</total_results>
//         <prospect>
//             <id>141844123</id>
//             <campaign_id>44516</campaign_id>
//             <salutation></salutation>
//             <first_name>Ashley</first_name>
//             <last_name>Brown</last_name>
//             <email>gusandash@gmail.com</email>
//             <password/>
//             <company>Century21</company>
//             <website></website>
//             <job_title></job_title>
//             <department></department>
//             <country>US</country>
//             <address_one>26000 Hwy ra</address_one>
//             <address_two></address_two>
//             <city>Pittsburg</city>
//             <state>MO</state>
//             <territory/>
//             <zip>65724</zip>
//             <phone>6604730260</phone>
//             <fax></fax>
//             <source>OrganicSiteReg</source>
//             <annual_revenue></annual_revenue>
//             <employees></employees>
//             <industry></industry>
//             <years_in_business/>
//             <comments/>
//             <notes/>
//             <score>0</score>
//             <grade/>
//             <last_activity_at>2019-08-05 00:02:27</last_activity_at>
//             <recent_interaction>Inactive since: Aug 5, 2019</recent_interaction>
//             <crm_lead_fid/>
//             <crm_contact_fid>0031I00000t97vqQAA</crm_contact_fid>
//             <crm_owner_fid>0051I000001STGrQAO</crm_owner_fid>
//             <crm_account_fid/>
//             <salesforce_fid>0031I00000t97vqQAA</salesforce_fid>
//             <crm_last_sync>2019-08-05 11:35:13</crm_last_sync>
//             <crm_url>https://epromos.my.salesforce.com/0031I00000t97vqQAA</crm_url>
//             <is_do_not_email></is_do_not_email>
//             <is_do_not_call></is_do_not_call>
//             <opted_out></opted_out>
//             <is_reviewed>1</is_reviewed>
//             <is_starred></is_starred>
//             <created_at>2019-08-05 00:02:27</created_at>
//             <updated_at>2019-08-06 07:10:02</updated_at>
//             <campaign>
//                 <id>44516</id>
//                 <name>ePromos Forms</name>
//             </campaign>
//             <assigned_to>
//                 <user>
//                     <id>13400982</id>
//                     <email>chandrell.scott@epromos.com</email>
//                     <first_name>Chandrell</first_name>
//                     <last_name>Scott</last_name>
//                     <job_title>eCommerce Consultant</job_title>
//                     <role>Sales Manager</role>
//                     <account>526761</account>
//                     <created_at>2018-04-09 12:55:47</created_at>
//                     <updated_at>2019-05-17 10:04:04</updated_at>
//                 </user>
//             </assigned_to>
//             <last_activity>
//                 <visitor_activity>
//                     <id>1939139289</id>
//                     <visitor_id>339700225</visitor_id>
//                     <prospect_id>141844123</prospect_id>
//                     <type>4</type>
//                     <type_name>Form Handler</type_name>
//                     <details>Sample Completed</details>
//                     <form_handler_id>4064</form_handler_id>
//                     <campaign>
//                         <id>44516</id>
//                         <name>ePromos Forms</name>
//                     </campaign>
//                     <created_at>2019-08-05 00:02:27</created_at>
//                 </visitor_activity>
//             </last_activity>
//             <Custom_12_Month_Revenue_Amount_ASI>0</Custom_12_Month_Revenue_Amount_ASI>
//             <Custom_12_Month_Rolling_Revenue_Amount>0</Custom_12_Month_Rolling_Revenue_Amount>
//             <Custom_2_Year_Revenue_ASI>0</Custom_2_Year_Revenue_ASI>
//             <Custom_2_Year_Revenue_Contact>0</Custom_2_Year_Revenue_Contact>
//             <Ad_Group>Promo Items</Ad_Group>
//             <Ad_ID>342752590934</Ad_ID>
//             <Ad_Network>Google search</Ad_Network>
//             <Contact_Full_Name>Ashley Brown</Contact_Full_Name>
//             <Contact_Owner_Role>Team Leah ecommerce Consultant</Contact_Owner_Role>
//             <Created_Date_Salesforce>2019-08-05</Created_Date_Salesforce>
//             <Device>Mobile devices with full browsers</Device>
//             <Do_Not_Call>false</Do_Not_Call>
//             <Do_Not_Direct_Mail>false</Do_Not_Direct_Mail>
//             <Do_Not_Email>false</Do_Not_Email>
//             <Do_Not_Pursue>false</Do_Not_Pursue>
//             <First_Assigned>2019-08-05T16:11:43.000Z</First_Assigned>
//             <Free_Email_Provider>false</Free_Email_Provider>
//             <Frequency>None</Frequency>
//             <GCLid>Cj0KCQjwhJrqBRDZARIsALhp1WSihhebsNQsdY8QClYrGzbBANc7uddfIQozg2z3I7vpZhSAaxVBZhwaAuCVEALw_wcB</GCLid>
//             <Keyword>promotional items</Keyword>
//             <Last_RFM_Date>2019-08-05</Last_RFM_Date>
//             <Lead_Full_Name>Ashley Brown</Lead_Full_Name>
//             <Lead_Status>Unqualified</Lead_Status>
//             <Lifetime_Revenue_ASI>0</Lifetime_Revenue_ASI>
//             <Lifetime_Revenue_Contact>0</Lifetime_Revenue_Contact>
//             <Monetary>None</Monetary>
//             <Nurtured>false</Nurtured>
//             <PPC_Code>QPPAAZ0827</PPC_Code>
//             <Recency>None</Recency>
//             <Sales_Channel>Dot Com</Sales_Channel>
//             <Segment>None</Segment>
//             <Site_ID>155730110</Site_ID>
//             <Submitted_Reg_All_NA>1</Submitted_Reg_All_NA>
//             <Submitted_Reg_All_News>0</Submitted_Reg_All_News>
//             <Submitted_Reg_All_SAC>1</Submitted_Reg_All_SAC>
//             <Submitted_Reg_All_SDAD>1</Submitted_Reg_All_SDAD>
//             <SugarCRM_Legacy_Lead_Status>Prospect</SugarCRM_Legacy_Lead_Status>
//             <UTM_Campaign>S_RLSA_NB__PromotionalItems_USA_All_Exact</UTM_Campaign>
//             <Wants_Catalog>false</Wants_Catalog>
//             <Wants_Monthly_Newsletter>false</Wants_Monthly_Newsletter>
//             <Wants_New_Arrival>true</Wants_New_Arrival>
//             <Wants_Newsletter>true</Wants_Newsletter>
//             <Wants_Sales_Email>true</Wants_Sales_Email>
//             <Wants_Specials>true</Wants_Specials>
//             <Website_Action>Sample</Website_Action>
//             <prospect_account_id>6911287</prospect_account_id>
//         </prospect>
