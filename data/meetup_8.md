---
title: "Meetup #8"
excerpt: This meetup is dedicated to explanation of Istio and OpenTracing.
event_date: 28.11.2018
event_time: "18:00"
event_place: 
  url: https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2607.068927407547!2d16.610185315840635!3d49.19925197932248!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x0%3A0x0!2zNDnCsDExJzU3LjMiTiAxNsKwMzYnNDQuNiJF!5e0!3m2!1sen!2scz!4v1528745893930
  label: U Palečka
event_meetup: zxvznpyxpblc
event_facebook: 1964860376939432
---

!!! We have a new space. Now we are in the dance room upstairs of Palecek. !!!

The doors open at 18:00
The program starts at 18:30

1) Milan Zamazal - Are virtual machines obsolete? 

In the past, hosting services used to run on bare metal servers or to utilize
emulated virtual machines.  Nowadays, they are being replaced with a container
platforms.  The current hype tells us that we want to make all our applications
and services cloud-native by containerizing them.

Container-like technologies are not new but their potential had been
underestimated for many years.  They finally got appropriate attention in Linux
in the last years and the industry is moving towards microservices running on
the container platforms.  What are the advantages of containers?  Do they miss
anything?  Can they replace virtual machines completely, sooner or
later?  Or do we still need emulated virtual machines and will we need them in
any foreseeable future?

No special knowledge is required to understand answers to those questions.

Milan is a senior software engineer, working for Red Hat in open virtualization division for past 4 years.
He is interested in free software, opensource, linux and emacs. I am not kidding, ask him about emacs :).

2) Petr Kotas - What is the next step for virtualization?

The containers are great. Unless you run a legacy workload, or you want to control the security tightly. When you have a scenario like this in your hands, you need more advanced tools. Lucky for us, these tools are already being built. Transformation to the cloud-native era is not a simple task especially when you have lots of old workloads that need modernization.

It turns out, containers miss some nice properties of virtual machines, such as isolated kernel. To fill the gap, hypervisor technology is being used as additional layer of isolation. Thus, resulting in more secure and isolated workloads.

You will learn how you can benefit from old school virtual machines in the new container native world. I will discuss existing technologies and provide an example of their use. You will also learn why you should or should not use hypervisor technologies when you run your workloads in the cloud.

Petr's journey to the cloud world started 9 year ago at VŠB University of Ostrava, where he was part of a core team working on a computational libraries for distibuted computing systems. He then continued his trip in
Seznam.cz where he was a lead developer of a cloud application used for search engine debuging. After a brief machine learning stop in his failed food recommendation startup he ended as a senior engineer at Red Hat.
He now works in Red Hat Virtualization team where he also works on Kubevirt project (https://github.com/kubevirt/kubevirt/).
