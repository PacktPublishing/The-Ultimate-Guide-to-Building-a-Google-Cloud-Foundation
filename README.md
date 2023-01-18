


# The Ultimate Guide to Building a Google Cloud Foundation

<a href="https://www.packtpub.com/cloud-networking/the-ultimate-guide-to-google-cloud-foundation?utm_source=github&utm_medium=repository&utm_campaign=9781803240855"><img src="https://static.packt-cdn.com/products/9781803240855/cover/smaller?" alt="Early Access" height="256px" align="right"></a>

This is the code repository for [The Ultimate Guide to Building a Google Cloud Foundation](https://www.packtpub.com/cloud-networking/the-ultimate-guide-to-google-cloud-foundation?utm_source=github&utm_medium=repository&utm_campaign=9781803240855), published by Packt.

**A one-on-one tutorial with one of Google’s top trainers**

## What is this book about?
From data ingestion and storage, through data processing and data analytics, to application hosting and even machine learning, whatever your IT infrastructural need, there's a good chance that Google Cloud has a service that can help. But instant, self-serve access to a virtually limitless pool of IT resources has its drawbacks. More and more organizations are running into cost overruns, security problems, and simple "why is this not working?" headaches. 

This book covers the following exciting features:
* Create an organizational resource hierarchy in Google Cloud
* Configure user access, permissions, and key Google Cloud Platform (GCP) security groups
* Construct well thought out, scalable, and secure virtual networks
* Stay informed about the latest logging and monitoring best practices
* Leverage Terraform infrastructure as code automation to eliminate toil
* Limit access with IAM policy bindings and organizational policies
* Implement Google's secure foundation blueprint

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1803240857) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" 
alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

A block of code is set as follows:
```
resource "google_tags_tag_value" "c_value" {
parent = "tagKeys/${google_tags_tag_key.c_key.name}"
short_name = "true"
description = "Project contains contracts."
}
```

Any command-line input or output is written as follows:
```
cd gcp-org
git checkout plan
```

**Following is what you need for this book:**
This book is for anyone looking to implement a secure foundational layer in Google Cloud, including cloud engineers, DevOps engineers, cloud security practitioners, developers, infrastructural management personnel, and other technical leads. A basic understanding of what the cloud is and how it works, as well as a strong desire to build out Google Cloud infrastructure the right way will help you make the most of this book. Knowledge of working in the terminal window from the command line will be beneficial.

A working knowledge of the terminal / command-line window is expected. Also, it would be helpful if you’ve played around in Google Cloud a little before reading the book.

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it](https://packt.link/FLbGs).

### Related products
* Professional Cloud Architect Google Cloud Certification Guide - Second Edition [[Packt]](https://www.packtpub.com/product/professional-cloud-architect-google-cloud-certification-guide-second-edition/9781801812290?utm_source=github&utm_medium=repository&utm_campaign=9781801812290) [[Amazon]](https://www.amazon.com/dp/1801812292)

* Architecting Google Cloud Solutions [[Packt]](https://www.packtpub.com/product/architecting-google-cloud-solutions/9781800563308?utm_source=github&utm_medium=repository&utm_campaign=9781800563308) [[Amazon]](https://www.amazon.com/dp/1800563302)

## Get to Know the Author
**Patrick Haggerty**
was never quite sure what he wanted to be when he grew up, so he decided he’d just try things until he figured it out. Thrown out of college at 20, he spent 4 years in the USMC learning responsibility (and to be a better apex predator). Out on a disability, he turned wrenches in a mechanic shop, worked tech support, studied Actuarial Science, and coded in more languages than he wants to remember. When a job asked him to run some internal training, he discovered a lifelong passion: helping people learn.
Patrick has worked as a professional trainer for 25+ years and spends most of his days working for ROI Training and Google, helping people learn to leverage Google Cloud.
### Download a free PDF

 <i>If you have already purchased a print or Kindle version of this book, you can get a DRM-free PDF version at no cost.<br>Simply click on the link to claim your free PDF.</i>
<p align="center"> <a href="https://packt.link/free-ebook/9781803240855">https://packt.link/free-ebook/9781803240855 </a> </p>